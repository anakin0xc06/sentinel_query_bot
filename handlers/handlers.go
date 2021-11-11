package handlers

import (
	"fmt"

	"github.com/anakin0xc06/sentinel_query_bot/db"
	"github.com/anakin0xc06/sentinel_query_bot/helpers"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"github.com/globalsign/mgo/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/telegram-bot-api.v4"
)

// MainHandler ...
func MainHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update,
	botDB *mongo.Collection) {

	if update.Message != nil && update.Message.IsCommand() {
		username := helpers.GetUserName(update)
		chatID := helpers.GetChatID(update)
		fmt.Println(username, chatID)
		command := update.Message.Command()

		switch command {
		case "start":
			StartHandler(bot, update, botDB)
		case "sentinel":
			AboutSentinel(bot, update)
		case "updates":
			Updates(bot, update)
		case "about":
			AboutBot(bot, update)
		case "subscribe":
			HandleSubscribe(bot, update, botDB)
		case "unsubscribe":
			HandleUnSubscribe(bot, update, botDB)
		case "help":
			HelpHandler(bot, update, botDB)
		case "restart":
			StartHandler(bot, update, botDB)

		default:
			text := "Command not available. see /help"
			helpers.SendMessage(bot, update, text, "html")
		}
	}

	if update.ChannelPost != nil && update.ChannelPost.IsCommand() {
		command := update.ChannelPost.Command()
		switch command {
		case "registerChannel":
			HandleRegisterChannel(bot, update, botDB)
		case "unregisterChannel":
			HandleUnRegisterChannel(bot, update, botDB)
		default:
			return
		}
	}
}

// StartHandler ...
func StartHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update,
	botDB *mongo.Collection) {

	username := helpers.GetUserName(update)
	text := fmt.Sprintf(templates.WelcomeGreetMsg, username) + "\n\n" + templates.SelectHelpMsg

	helpers.SendMessage(bot, update, text, tgbotapi.ModeMarkdown)
	chatID := helpers.GetChatID(update)
	userUpdate := bson.M{"$set": bson.M{"username": username, "chatid": chatID, "status": "HOME"}}
	db.UpdateUser(botDB, chatID, userUpdate)
}

// HelpHandler ...
func HelpHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, botDB *mongo.Collection) {
	text := templates.HelpMsg
	helpers.SendMessage(bot, update, text, tgbotapi.ModeHTML)
	db.UpdateStatus(botDB, helpers.GetUserName(update),
		helpers.GetChatID(update), "HELP")
}

func HandleSubscribe(bot *tgbotapi.BotAPI, update tgbotapi.Update, botDB *mongo.Collection) {
	username := helpers.GetUserName(update)
	update.Message.Chat.IsGroup()
	botDetails, _ := bot.GetMe()
	text := fmt.Sprintf(templates.SubscribeMsg, botDetails.UserName)
	chatID := helpers.GetChatID(update)
	userUpdate := bson.M{"$set": bson.M{"username": username, "chatid": chatID, "status": "HOME", "subscribed": true}}
	db.UpdateUser(botDB, chatID, userUpdate)
	helpers.SendMessage(bot, update, text, tgbotapi.ModeHTML)
}

func HandleUnSubscribe(bot *tgbotapi.BotAPI, update tgbotapi.Update, botDB *mongo.Collection) {
	username := helpers.GetUserName(update)
	botDetails, _ := bot.GetMe()
	text := fmt.Sprintf(templates.UnSubscribeMsg, botDetails.UserName)
	chatID := helpers.GetChatID(update)
	userUpdate := bson.M{"$set": bson.M{"username": username, "chatid": chatID, "status": "HOME", "subscribed": false}}
	db.UpdateUser(botDB, chatID, userUpdate)
	helpers.SendMessage(bot, update, text, tgbotapi.ModeHTML)
}

func HandleRegisterChannel(bot *tgbotapi.BotAPI,
	update tgbotapi.Update, botDB *mongo.Collection) {
	chatID := update.ChannelPost.Chat.ID
	username := update.ChannelPost.Chat.UserName
	userUpdate := bson.M{"$set": bson.M{"username": username, "chatid": chatID, "subscribed": true}}
	db.UpdateUser(botDB, chatID, userUpdate)
	text := templates.ChannelRegisteredMsg
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "html"
	bot.Send(msg)
}

func HandleUnRegisterChannel(bot *tgbotapi.BotAPI,
	update tgbotapi.Update, botDB *mongo.Collection) {
	chatID := update.ChannelPost.Chat.ID
	username := update.ChannelPost.Chat.UserName
	userUpdate := bson.M{"$set": bson.M{"username": username, "chatid": chatID, "subscribed": false}}
	db.UpdateUser(botDB, chatID, userUpdate)
	text := templates.ChannelUnRegisteredMsg
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "html"
	bot.Send(msg)

}
