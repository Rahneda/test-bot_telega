# Simple Telegram Bot

A simple Telegram bot written in Go that responds with "hi" to any incoming message.

## Prerequisites

- Go 1.20 or higher
- A Telegram Bot Token (get it from [@BotFather](https://t.me/BotFather))

## Installation

1. Clone this repository
2. Install dependencies:
   ```
   go mod download
   ```

## Usage

1. Set your Telegram bot token as an environment variable:

   ```
   export TELEGRAM_TOKEN=your_telegram_bot_token
   ```

2. Run the bot:

   ```
   go run main.go
   ```

3. Send any message to your bot, and it will respond with "hi".

## How it works

The bot uses the [gotgbot](https://github.com/PaulSonOfLars/gotgbot) library to interact with the Telegram Bot API. It sets up a message handler that responds to all incoming messages with the text "hi". 