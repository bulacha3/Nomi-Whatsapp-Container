# Base image
FROM golang:1.23.1-alpine

# Instalar dependências necessárias
RUN apk add --no-cache git build-base

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos necessários para o build
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o projeto para dentro do container
COPY . .

# Copiar explicitamente o store.db
COPY store.db /app/store.db

# Garantir permissões adequadas
RUN chmod 777 /app/store.db

# Compilar o aplicativo Go
RUN go build -o nomi-whatsapp ./cmd/generic/main.go

# Expor a porta do aplicativo
EXPOSE 8080

# Comando de inicialização
CMD ["/app/nomi-whatsapp"]




