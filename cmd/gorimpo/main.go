package main

import (
	"log/slog"
	"os"
	"time"

	"github.com/LXSCA7/gorimpo/internal/adapters/notifier"
	"github.com/LXSCA7/gorimpo/internal/adapters/repository"
	"github.com/LXSCA7/gorimpo/internal/adapters/scraper"
	"github.com/LXSCA7/gorimpo/internal/config"
	"github.com/LXSCA7/gorimpo/internal/core/services"
	"github.com/joho/godotenv"
	"github.com/lmittmann/tint"
)

var Version = "dev"

func main() {
	logger := slog.New(tint.NewHandler(os.Stdout, &tint.Options{
		Level:      slog.LevelDebug,
		TimeFormat: time.TimeOnly,
	}))
	slog.SetDefault(logger)
	routes := make(map[string]string)

	cfg, err := config.Load("./config.yaml")
	if err != nil {
		panic(err)
	}

	for _, cat := range cfg.Categories {
		if !cfg.App.UseTopics {
			routes[cat] = "0"
		} else {
			if cat == "nintendo" {
				routes[cat] = "3"
			} else {
				routes[cat] = "0"
			}
		}
	}
	_ = godotenv.Load()

	token := os.Getenv("TELEGRAM_TOKEN")
	chatID := os.Getenv("TELEGRAM_CHAT_ID")

	if token == "" || chatID == "" {
		logger.Error("missing TELEGRAM_TOKEN or TELEGRAM_CHAT_ID")
		os.Exit(1)
	}

	telegram := notifier.NewTelegram(token, chatID)
	telegram.SetRoutes(routes)

	olxScraper := scraper.NewOLX()
	if err := os.MkdirAll("data", os.ModePerm); err != nil {
		logger.Error("Erro ao criar pasta data", "erro", err)
		os.Exit(1)
	}

	repo, err := repository.NewSQLite("data/gorimpo.db")
	if err != nil {
		logger.Error("Erro ao iniciar o banco de dados", "erro", err)
		os.Exit(1)
	}
	gorimpoSvc := services.NewGorimpoService(olxScraper, repo, telegram, cfg)

	slog.Info("🚀 GOrimpo starting...", slog.String("version", Version))
	gorimpoSvc.Start(Version)
	logger.Info("👋 Sistema encerrado.")
}
