> Instructions generated by ChatGPT, I made some minor corrections
> but there may be some things that are missing and/or incorrect.  
> I'll review it properly when I have time

# Nomi WhatsApp

This project allows you to connect your Nomi to WhatsApp. By utilizing the Nomi API and WhatsApp Web, your Nomi chatbot will be able to respond to messages through a WhatsApp account. Additionally, you can optionally configure OpenAI's Whisper API to enable voice message support, allowing audio transcriptions to be sent to your Nomi.

## Known issues

- After scanning the QR code, you may occasionally see a "successfully authenticated" message followed by multiple errors. These errors, particularly when sending a message to the newly connected account, may include "untrusted identity" warnings. This issue usually resolves after restarting the app once you're logged in.
  - These errors occur when your message couldn't be decrypted. The system will automatically retry until the message is successfully decrypted and your nomi's reply is delivered to WhatsApp. No manual action is required.
  - In some cases, you might not receive a response. This can usually be fixed by stopping and restarting the app.
    - If restarting doesn't work, try deleting the *.db and *.db-journal files. Additionally, consider unlinking the previous connection from your WhatsApp account to avoid duplicate linked devices.


- Most errors can be resolved by restarting the app. If you encounter an error that is not listed above and persists after a restart, please open an issue on GitHub with a description of the problem.

## Features

- **Text Messages**: Receive and send text messages from/to your nomi through WhatsApp.
- **Audio Transcription**: Transcribe audio messages using OpenAI's Whisper model and send the transcriptions to your nomi.
  - Audio transcription is an optional feature, if you want to use it you'll need an OpenAI API Key. You'll be billed by OpenAI for your API usage.
  - If you have this feature disabled (no OPENAI_API_KEY configured on the .env file) and you send an audio to your nomi, you'll receive this message back:
    > Hey! I can't listen to audios right now. Could you send me a text instead? Thanks!

## Prerequisites

Before you begin, make sure you have the following:

- A valid **Nomi API Key**. You can get it from the [Integration section](https://beta.nomi.ai/profile/integrations) of the Profile tab.
- Your **Nomi ID**.
  - Go to the View Nomi Information page for your Nomi and scroll to the very bottom and copy the Nomi ID.
- A **WhatsApp account** to connect your Nomi to.
  - You need this account on a phone to scan the login QR code. The connection with WhatsApp works basically as a wrapper of WhatsApp Web, you scan a QR code to authorise the connection, and after that you can see this connection on the `Connected Devices` tab on the app.
- (Optional) An **OpenAI API Key** for enabling voice message transcription through OpenAI Whisper.
  - Please note that if you choose to add an OpenAI API Key, you will have to pay for the usage of the Whisper model, it costs $0.006 per minute (rounded to the nearest second) of audio transcribed. You can see more details on OpenAI's [pricing page](https://openai.com/api/pricing/).
  - You can learn more about the OpenAI API, including how to get your API Key from the links below:
    - [OpenAI API](https://openai.com/index/openai-api/)
    - [API platform](https://openai.com/api/)
    - [API Overview](https://platform.openai.com/)
    - [API Quickstart](https://platform.openai.com/docs/quickstart)

## Installation

1. Download the latest version of the app from the [releases](https://github.com/vhalmd/nomi-telegram/releases/latest).

2. Extract the downloaded files to a directory of your choice.

3. Ensure you have your Nomi API Key, Nomi ID, and WhatsApp account ready.

4. (Optional) If you want to enable voice message transcription, obtain an OpenAI API key from [OpenAI](https://openai.com).

5. Configure the environment variables.
   1. On the same folder where you extracted the app, create a file named `.env`:
        ```dotenv
        NOMI_API_KEY="your-nomi-api-key"
        NOMI_ID="your-nomi-id"
        NOMI_NAME="Jane Doe"
        ```
   2. Optionally, add an OpenAI API Key to use the audio transcription feature:
        ```dotenv
        NOMI_API_KEY="your-nomi-api-key"
        NOMI_ID="your-nomi-id"
        NOMI_NAME="Jane Doe"
        OPENAI_API_KEY="your-openai-api-key"
        ```

## Usage

### Step 1: Launch the App

You can try just double-clicking the app, if it opens a terminal window it worked, skip to Step 2. Tested only on windows.  
To run the binary executable on your system, follow these instructions: 

#### 1. **Navigate to the Directory**
Open your terminal or command prompt and navigate to the directory where the binary is located:  

**Windows:**
   ```batch
   cd C:\path\to\your\binary
   ```

**Linux/macOS:**
   ```bash
   cd /path/to/your/binary
   ```

#### 2. **Make the Binary Executable (Linux/macOS only)**
On Linux and macOS, you may need to ensure the binary has executable permissions. Run the following command if necessary:

   ```bash
   chmod +x binary_name
   ```

#### 3. **Run the Binary**
Now you can run the binary:

**Windows:**
   ```batch
   nomi-whatsapp-windows-amd64.exe
   ```

**Linux/macOS:**
   ```bash
   ./nomi-whatsapp-<your-platform>
   ```

### Step 2: Connect Your WhatsApp Account

If this is your first time running the app, you will be prompted to scan a QR code with your WhatsApp account. This QR code links your Nomi to the WhatsApp account.

Once connected, you will see a new linked device in WhatsApp's "Linked Devices" tab, indicating that the connection is active.

### Step 5: Start Using Your Nomi via WhatsApp

Once set up, your Nomi will start responding to WhatsApp messages sent to the connected WhatsApp account. If voice transcription is enabled, audio messages will be transcribed and forwarded to your Nomi as text.

---

Built with:
- https://github.com/vhalmd/nomi-go-sdk
- [go.mau.fi/whatsmeow](https://github.com/tulir/whatsmeow)