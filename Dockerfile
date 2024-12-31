# Usar uma imagem base do Go mais recente
FROM golang:1.23.1-alpine

# Definir o diretório de trabalho
WORKDIR /app

# Copiar os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Baixar as dependências
RUN go mod download

# Copiar todo o código para o container
COPY . .

# Compilar o aplicativo
RUN go build -o nomi-whatsapp ./cmd/generic/main.go

# Expor a porta que o aplicativo irá usar
EXPOSE 8080

# Configurar variáveis de ambiente
ENV PORT=8080

# Comando para rodar o aplicativo
CMD ["sh", "-c", "cp store.db /tmp/store.db && ./nomi-whatsapp"]
