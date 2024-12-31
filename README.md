# Nomi-Whatsapp-Container

# Nomi WhatsApp Bot

Este Bot é uma melhoria do bot do vhalmd que pode ser visto aqui: https://github.com/vhalmd/nomi-whatsapp

ele transforma o bot em container permitindo vc a rodar localmente ou na nuvem como no Google Cloud Run.

## Pré-requisitos

- Docker
- Conta no Google Cloud (para usar o Cloud Run)

## Passos para Configuração

1. Configure as variáveis de ambiente no Cloud Run:
   - `NOMI_API_KEY`
   - `NOMI_ID`
   - `NOMI_NAME`
   - `OPENAI_API_KEY`

2. Construa e rode o Docker localmente:
   ```bash
   docker build -t nomi-whatsapp .
   docker run -p 8080:8080 --env-file .env nomi-whatsapp
