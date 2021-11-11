package handlers

import (
	"github.com/anakin0xc06/sentinel_query_bot/helpers"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"gopkg.in/telegram-bot-api.v4"
)

// AboutSentinel ...
func AboutSentinel(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := templates.AboutSentinel
	helpers.SendMessage(bot, update, text, "html")
}

// AboutBot ...
func AboutBot(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := templates.AboutBot
	helpers.SendMessage(bot, update, text, "html")
}