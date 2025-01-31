package discord

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DiscordWebhook struct {
	Content string `json:"content"`
}

func SendMessage(webhookURL, message string) error {
	payload := DiscordWebhook{
		Content: message,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("JSON marshal error: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("HTTP post error: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("Discord API error: %s", resp.Status)
	}

	return nil
}
