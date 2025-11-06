# Nomi-Whatsapp-Container

## Nomi WhatsApp Bot

Este Bot é uma melhoria do bot do [vhalmd](https://github.com/vhalmd/nomi-whatsapp), que agora está adaptado para funcionar como um container. Ele permite a execução localmente ou em nuvens como o Google Cloud Run.

---

## Pré-requisitos

- **Docker** instalado e configurado em sua máquina.
- Uma conta no **Google Cloud** com o Cloud Run habilitado.
- Acesso às seguintes chaves e informações:
  - `NOMI_API_KEY`
  - `NOMI_ID`
  - `NOMI_NAME`
  - `OPENAI_API_KEY`
- **Arquivo `store.db`** (necessário para conectar ao WhatsApp). Caso você não tenha, siga os passos na seção "Como Gerar o Arquivo `store.db`".

---
#### Passo 1: Clone o Repositório
```bash
git clone https://github.com/bulacha3/Nomi-Whatsapp-Container.git
cd Nomi-Whatsapp-Container
```
## Como Gerar o Arquivo `store.db`

1. **Compile um executável atualizado**:
   - Instale o Go 1.24.x (ou superior compatível).
   - Dentro deste repositório, execute `./build.sh` para gerar os binários em `bin/`.
   - O script cria versões para Linux, macOS e Windows; escolha a que corresponde ao seu sistema.
2. **Execute o binário a partir de um terminal**:
   - Mantenha o arquivo `.env` na mesma pasta para que o `godotenv` carregue as variáveis automaticamente.
   - No Windows, abra o PowerShell ou Prompt de Comando e execute `bin\nomi-whatsapp-windows-amd64.exe`.
   - Em Linux/macOS, execute `./bin/nomi-whatsapp-linux-amd64` (ou a variante da sua arquitetura).
   - O programa roda inteiramente no terminal e não abre mais nenhuma página web.
3. **Escaneie o código QR exibido no terminal**:
   - Assim que conectar ao WhatsApp, um QR code em ASCII aparecerá na tela.
   - Escaneie com o aplicativo WhatsApp no celular. A mensagem `QR code accepted!` confirma o pareamento.
   - Se o QR expirar, aguarde: um novo será impresso automaticamente.
4. **Salve o arquivo gerado**:
   - Após a autenticação, o arquivo `store.db` é criado na pasta onde você executou o binário.
   - Faça backup desse arquivo; ele será reutilizado nos próximos deploys.

---

## Configuração e Deploy

###Configuração Local

  #### 1: Configure o Arquivo `.env`
  Crie um arquivo `.env` na raiz do projeto com as seguintes variáveis:

  ```env
  NOMI_API_KEY=Chave_API_nomi
  NOMI_ID=nomi_id
  NOMI_NAME=nome_do_nomi
  OPENAI_API_KEY=chave_api_openai (se quiser caso contrario nao mencione no arquivo .env)
  ```

#### Passo 1: Construa e Rode o Docker Localmente
  ```bash
  docker build -t nomi-whatsapp .
  docker run -p 8080:8080 --env-file .env nomi-whatsapp
  ```

  ---

### 2. Deploy no Google Cloud Run

  #### Passo 2.1: Faça Login no Google Cloud
  Certifique-se de estar autenticado no Google Cloud:
  ```bash
    gcloud auth login
  ```

  #### Passo 2.2: Configure o Projeto do Google Cloud
  ```bash
    gcloud config set project [SEU_ID_DO_PROJETO]
  ```

  #### Passo 3: Habilite APIs Necessárias
  Habilite o serviço do Cloud Run (caso ainda não esteja ativado):
  ```bash
      gcloud services enable run.googleapis.com
  ```

  #### Passo 4: Suba a Imagem para o Google Container Registry
  ```bash
      docker tag nomi-whatsapp gcr.io/[SEU_ID_DO_PROJETO]/nomi-whatsapp

    docker push gcr.io/[SEU_ID_DO_PROJETO]/nomi-whatsapp
  ```

  #### Passo 5: Faça o Deploy no Cloud Run
   ##### lembre-se que estes comandos podem variar na estrutura devido ao seu terminal 
  ```bash
    gcloud run deploy nomi-whatsapp \
      --image gcr.io/[SEU_ID_DO_PROJETO]/nomi-whatsapp \
      --region southamerica-east1 \
      --platform managed \
      --allow-unauthenticated \
      --set-env-vars NOMI_API_KEY=sua_chave_api_nomi,NOMI_ID=seu_id_nomi,NOMI_NAME=seu_nome_nomi,OPENAI_API_KEY=sua_chave_api_openai
  ```

  ---

## Debug e Logs

### Ver Logs no Cloud Run
Você pode verificar os logs do serviço para identificar possíveis erros:
```bash
gcloud run services logs read nomi-whatsapp
```

### Erros Comuns
- **Erro de Permissão no Store.db**: Certifique-se de que o arquivo `store.db` está corretamente referenciado no Dockerfile e incluído na imagem Docker.
- **Timeout de Inicialização**: Verifique se todas as variáveis de ambiente estão corretamente configuradas e o serviço está escutando na porta 8080.

---

## Contribuição
Contribuições são bem-vindas! Por favor, abra um Pull Request ou envie sugestões na aba de Issues.

---

## Licença
Este projeto está sob a licença MIT. Consulte o arquivo `LICENSE` para mais detalhes.

