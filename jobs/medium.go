package jobs

import (
	"fmt"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"strings"
	"time"

	"github.com/jasonlvhit/gocron"
	"github.com/mmcdole/gofeed"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"go.mongodb.org/mongo-driver/mongo"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

var lastPublishTime *time.Time
var feedParser = gofeed.NewParser()

var MediumFeedURL = fmt.Sprintf("https://medium.com/feed/@%s",config.MediumHandle)

// CheckForNewPublication ...
func CheckForNewPublication(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	feed, err := feedParser.ParseURL(MediumFeedURL)
	if err != nil {
		return
	}

	post := feed.Items[0]
	if post.PublishedParsed.String() != lastPublishTime.String() {
		lastPublishTime = post.PublishedParsed
		fmt.Println("New Publication: ", post.Title)
		link := strings.Split(post.Link, "?")[0]
		txt := fmt.Sprintf(templates.MediumPost, post.Title, link)
		users := GetAllChatIDs(db)
		broadcastMediumPost(bot, users, txt)
	}
}

func broadcastMediumPost(bot *tgbotapi.BotAPI, chatIDs []int64, text string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, text)
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}
	return
}
func GetLatestPublicationTime() *time.Time {
	feed, err := feedParser.ParseURL(MediumFeedURL)
	if err != nil {
		fmt.Println(err)
	}
	return feed.Items[0].PublishedParsed
}

// MediumPostScheduler ...
func MediumPostScheduler(bot *tgbotapi.BotAPI, db *mongo.Collection) {
	fmt.Println("[Medium Job]: successfully started medium post notifier job")
	lastPublishTime = GetLatestPublicationTime()
	s := gocron.NewScheduler()
	s.Every(10).Seconds().Do(CheckForNewPublication, bot, db)
	<-s.Start()
}
