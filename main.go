package main

import (
	"flag"
	"fmt"
	"os"

	"discord-message-cli/discord" // discordパッケージをimport
)

func main() {
	webhookURL := flag.String("url", "", "Discord Webhook URL")
	message := flag.String("message", "Hello, Discord from CLI!", "Message to send")
	flag.Parse()

	if *webhookURL == "" {
		fmt.Println("Error: --url is required")
		flag.Usage()
		os.Exit(1)
	}

	if err := discord.SendMessage(*webhookURL, *message); err != nil { // discord.SendMessageを使用
		fmt.Println("Error sending message:", err)
		os.Exit(1)
	}

	fmt.Println("Message sent successfully!")
}
