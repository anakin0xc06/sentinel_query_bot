package main

import (
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"log"
	"os"

	"github.com/anakin0xc06/sentinel_query_bot/db"
	"github.com/anakin0xc06/sentinel_query_bot/handlers"
	"github.com/anakin0xc06/sentinel_query_bot/jobs"
	"github.com/fatih/color"
	"gopkg.in/telegram-bot-api.v4"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_API_KEY"))
	if err != nil {
		log.Fatalf("Error in instantiating the bot: %v", err)
	}

	mongoDb := db.NewDB()
	botDB := mongoDb.Database("sentinel_query_bot").Collection("users")

	if len(config.TwitterHandle) > 0 && len(config.TwitterID) > 0 {
		stream := jobs.TwitterConfig()
		go jobs.ListenAndNotifyTweets(bot, stream, botDB)
	}
	if len(config.MediumHandle) > 0 {
		go jobs.MediumPostScheduler(bot, botDB)
	}
	if len(config.RedditHandle) > 0 {
		go jobs.RedditPostScheduler(bot, botDB)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		color.Red("Error while receiving messages: %s", err)
		return
	}
	color.Green("Started %s successfully", bot.Self.UserName)

	for update := range updates {
		handlers.MainHandler(bot, update, botDB)
	}
}
