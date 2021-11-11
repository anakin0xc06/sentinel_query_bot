package handlers

import (
	"fmt"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"log"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/mmcdole/gofeed"
	"github.com/anakin0xc06/sentinel_query_bot/helpers"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"gopkg.in/telegram-bot-api.v4"
)

var fp = gofeed.NewParser()
var MediumFeedURL = fmt.Sprintf("https://medium.com/feed/@%s",config.MediumHandle)
var RedditFeedURL = fmt.Sprintf("https://www.reddit.com/%s.rss", config.RedditHandle)

// Updtes

func Updates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cmd := update.Message.Command()
	args := update.Message.CommandArguments()
	fmt.Println(cmd, "|", args)
}
// MediumUpdates ...
func MediumUpdates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	feed, err := fp.ParseURL(MediumFeedURL)
	if err != nil {
		log.Println(err)
		return
	}
	posts := feed.Items
	for idx, post := range posts {
		if idx > 2 {
			break
		}
		link := strings.Split(post.Link, "?")[0]
		text := fmt.Sprintf(templates.MediumFormat, post.Title, link)
		helpers.SendMessage(bot, update, text, "html")
	}
	return
}

// RedditUpdates ...
func RedditUpdates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	feed, err := fp.ParseURL(RedditFeedURL)
	if err != nil {
		log.Println(err)
		return
	}
	for idx, post := range feed.Items {
		if idx > 2 {
			break
		}
		text := fmt.Sprintf(templates.RedditFormat, post.Updated, post.Title, post.Link)
		helpers.SendMessage(bot, update, text, "html")
	}
	return
}

// TwitterUpdates ...
func TwitterUpdates(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	cfg := oauth1.NewConfig(config.TwitterConsumerKey, config.TwitterConsumerSecret)
	token := oauth1.NewToken(config.TwitterAccessToken, config.TwitterAccessTokenSecret)

	httpClient := cfg.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	tweets, _, err := client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: config.TwitterHandle,
		Count:      3,
		TweetMode:  "extended",
	})
	if err != nil {
		log.Println("Error while getting tweets: ", err)
	}
	for _, tweet := range tweets {
		text := fmt.Sprintf(templates.TweetFormat, tweet.CreatedAt, tweet.FullText)
		helpers.SendMessage(bot, update, text, "html")
	}
	return
}
