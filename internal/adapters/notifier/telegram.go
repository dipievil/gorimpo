package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/LXSCA7/gorimpo/internal/core/domain"
)

type TelegramAdapter struct {
	Token  string
	ChatID string
	ApiURL string
}

func NewTelegram(token, chatID string) *TelegramAdapter {
	return &TelegramAdapter{
		Token:  token,
		ChatID: chatID,
		ApiURL: fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token),
	}
}

func (t *TelegramAdapter) SendText(message string) error {
	payload := map[string]string{
		"chat_id":    t.ChatID,
		"text":       message,
		"parse_mode": "HTML",
	}
	return t.doRequest(payload)
}

func (t *TelegramAdapter) Send(offer domain.Offer) error {
	msg := fmt.Sprintf(
		"🚨 <b>NOVO ACHADO NO %s!</b>\n\n🎮 <b>%s</b>\n💰 Preço: <b>R$ %.2f</b>\n\n🔗 <a href=\"%s\">Ver Anúncio</a>",
		offer.Source, offer.Title, offer.Price, offer.Link,
	)

	return t.SendText(msg)
}

func (t *TelegramAdapter) doRequest(payload map[string]string) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(t.ApiURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		slog.Error("erro: %s", resp.Body)
		return fmt.Errorf("erro na api do telegram: status %d", resp.StatusCode)
	}

	return nil
}
