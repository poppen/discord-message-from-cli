package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"discord-message-cli/discord" // discordパッケージをimport
)

type Config struct {
	WebhookURL string `json:"webhook_url"`
}

func loadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func main() {
	webhookURL := flag.String("url", "", "Discord Webhook URL")
	message := flag.String("message", "Hello, Discord from CLI!", "Message to send")
	configPath := flag.String("config", "config.json", "Path to config file") // configファイルのパスを指定するflagを追加
	flag.Parse()

	cfg, err := loadConfig(*configPath)    // configファイルの読み込み
	if err != nil && !os.IsNotExist(err) { // ファイルが存在しない場合はエラーを無視
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	if *webhookURL == "" {
		if cfg != nil && cfg.WebhookURL != "" {
			webhookURL = &cfg.WebhookURL // configファイルからwebhook URLを取得
		} else {
			fmt.Println("Error: --url is required or webhook_url in config.json")
			flag.Usage()
			os.Exit(1)
		}
	}

	if err := discord.SendMessage(*webhookURL, *message); err != nil { // discord.SendMessageを使用
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully!")
}
