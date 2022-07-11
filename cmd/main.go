package main

import (
	"github.com/perfectogo/template_for_telegram_bot/bot"
	"github.com/perfectogo/template_for_telegram_bot/config"
)

func main() {
	cfg := config.Load()
	bot.Bot(cfg.TelegramBotApiToken)

}
