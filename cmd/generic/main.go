package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"github.com/mdp/qrterminal/v3"
	"github.com/vhalmd/nomi-whatsapp/internal/whatsapp"
	waLog "go.mau.fi/whatsmeow/util/log"
	_ "modernc.org/sqlite"
)

func main() {
	// Configuração das variáveis de ambiente
	nomiApiKey := os.Getenv("NOMI_API_KEY")
	nomiID := os.Getenv("NOMI_ID")
	nomiName := os.Getenv("NOMI_NAME")
	openAIToken := os.Getenv("OPENAI_API_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Verifica variáveis obrigatórias
	if nomiApiKey == "" || nomiID == "" || nomiName == "" {
		log.Fatal("As variáveis NOMI_API_KEY, NOMI_ID e NOMI_NAME são obrigatórias!")
	}

	// Inicia o servidor HTTP em uma goroutine
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Nomi WhatsApp está rodando!"))
		})
		log.Printf("Servidor rodando na porta %s", port)
		log.Fatal(http.ListenAndServe(":"+port, nil))
	}()

	// Configuração do cliente WhatsApp
	cfg := whatsapp.Config{
		NomiAPIKey: nomiApiKey,
		NomiID:     nomiID,
		NomiName:   nomiName,
		OpenAIKey:  openAIToken,
	}

	clientLog := waLog.Stdout("CLIENT", "INFO", true)
	client := whatsapp.NewClient(cfg, clientLog)
	client.Whatsapp.AddEventHandler(client.EventHandler)

	// Login ou reconexão
	if client.Whatsapp.Store.ID == nil {
		qrChan, _ := client.Whatsapp.GetQRChannel(context.Background())
		err := client.Whatsapp.Connect()
		if err != nil {
			log.Fatalf("Erro ao conectar ao WhatsApp: %v", err)
		}

		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else {
				log.Printf("Evento de login: %s", evt.Event)
			}
		}
	} else {
		err := client.Whatsapp.Connect()
		if err != nil {
			log.Fatalf("Erro ao conectar ao WhatsApp: %v", err)
		}
	}

	// Finalização do programa
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Desconectando o cliente do WhatsApp...")
		client.Whatsapp.Disconnect()
		os.Exit(0)
	}()

	// Aguarda sinal para finalizar
	select {}
}
