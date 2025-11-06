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

	err := http.ListenAndServe(":5555", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *API) GetQR(w http.ResponseWriter, r *http.Request) {
	var status string
	if a.Client.Whatsapp.Store.ID == nil {
		status = "disconnected"
	} else {
		status = "connected"
	}

	response := map[string]string{
		"status": status,
		"qr":     a.Client.CurrentQRCode(),
	}
	data, _ := json.Marshal(response)
	w.WriteHeader(202)
	_, _ = w.Write(data)
}
