package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vhalmd/nomi-whatsapp/internal/whatsapp"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	nomiAPIKey := os.Getenv("NOMI_API_KEY")
	nomiID := os.Getenv("NOMI_ID")
	nomiName := os.Getenv("NOMI_NAME")
	openAIKey := os.Getenv("OPENAI_API_KEY")

	cfg := whatsapp.Config{
		NomiAPIKey: nomiAPIKey,
		NomiID:     nomiID,
		NomiName:   nomiName,
		OpenAIKey:  openAIKey,
	}

	clientLog := waLog.Stdout("CLIENT", "INFO", true)
	client := whatsapp.NewClient(cfg, clientLog)
	client.Whatsapp.EnableAutoReconnect = true
	client.Whatsapp.AddEventHandler(client.EventHandler)

	go client.ListenQR()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh
	fmt.Println("Shutting down…")
	client.Whatsapp.Disconnect()
}
