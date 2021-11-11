package jobs

import (
	"fmt"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/mmcdole/gofeed"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/telegram-bot-api.v4"
)

var lastUpdateTime *time.Time
var fp = gofeed.NewParser()

var RedditFeedURL = fmt.Sprintf("https://www.reddit.com/%s.rss", config.RedditHandle)

// CheckForNewPost ...
func CheckForNewPost(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	feed, err := fp.ParseURL(RedditFeedURL)
	if err != nil {
		return
	}
	post := feed.Items[0]
	if post.UpdatedParsed.String() != lastUpdateTime.String() {
		lastUpdateTime = post.UpdatedParsed
		fmt.Println("New Reddit Post: ", post.Title)
		fmt.Println(post.Link)
		text := fmt.Sprintf(templates.RedditPost, post.Author.Name, post.Title, post.Link)
		users := GetAllChatIDs(db)
		fmt.Println(users)
		go broadcastRedditPost(bot, users, text)
	}
}

func broadcastRedditPost(bot *tgbotapi.BotAPI, chatIDs []int64, text string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, text)
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}
}

func GetLatestPostTime() *time.Time {
	feed, err := fp.ParseURL(RedditFeedURL)
	if err != nil {
		fmt.Println(err)
	}
	return feed.Items[0].UpdatedParsed
}

// RedditPostScheduler ...
func RedditPostScheduler(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	fmt.Println("[Reddit Job]: successfully started reddit post notifier job")
	lastUpdateTime = GetLatestPostTime()
	s := gocron.NewScheduler()
	s.Every(10).Seconds().Do(CheckForNewPost, bot, db)
	<-s.Start()
}
