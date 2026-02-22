package services

import (
	"fmt"
	"log/slog"

	"github.com/LXSCA7/gorimpo/internal/core/ports"
)

type SystemService struct {
	repo          ports.SystemRepository
	notifier      ports.Notifier
	configManager ports.ConfigManager
}

func NewSystemService(r ports.SystemRepository, n ports.Notifier, c ports.ConfigManager) *SystemService {
	return &SystemService{
		repo:          r,
		notifier:      n,
		configManager: c,
	}
}

func (s *SystemService) Setup(currentVersion string) map[string]string {
	routes := s.setupRoutes()
	s.notifier.SetRoutes(routes)
	s.checkVersion(currentVersion)
	return routes
}

func (s *SystemService) checkVersion(currentVersion string) {
	lastVersion := s.repo.GetLastVersion()

	if lastVersion != "" && lastVersion != currentVersion {
		slog.Info("🎉 Atualização detectada!", "old", lastVersion, "new", currentVersion)

		changelogMsg := fmt.Sprintf(
			"🚀 <b>GOrimpo Atualizado com Sucesso!</b>\n\n"+
				"De: <code>%s</code>\nPara: <code>%s</code>\n\n"+
				"🔗 <a href=\"https://github.com/LXSCA7/gorimpo/releases\">Ver Changelog</a>",
			lastVersion, currentVersion,
		)
		_ = s.notifier.SendText(changelogMsg, "system")
	}

	_ = s.repo.SetCurrentVersion(currentVersion)
}

func (s *SystemService) setupRoutes() map[string]string {
	config := s.configManager.Get()
	slog.Info("🗺️ Configurando rotas do sistema...")
	routes := make(map[string]string)

	cats := []string{"system"}
	cats = append(cats, config.Categories...)
	for _, cat := range cats {
		if !config.App.UseTopics {
			routes[cat] = "0"
			continue
		}

		destID := s.repo.GetRoute(cat)
		if destID == "" {
			slog.Info("✨ Criando novo tópico no Telegram...", "categoria", cat)

			newID, err := s.notifier.CreateCategory(cat)
			if err != nil {
				slog.Error("Erro ao criar tópico, jogando pro Geral", "erro", err)
				newID = "0"
			} else {
				_ = s.repo.SaveRoute(cat, newID)
			}
			destID = newID
		}
		routes[cat] = destID
	}

	return routes
}
