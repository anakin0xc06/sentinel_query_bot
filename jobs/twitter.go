package jobs

import (
	"context"
	"fmt"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"log"
	"strings"

	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"github.com/anakin0xc06/sentinel_query_bot/types"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/telegram-bot-api.v4"
)

// Credentials ...
type Credentials struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

func getClient(creds *Credentials) (*twitter.Client, error) {
	config := oauth1.NewConfig(creds.ConsumerKey, creds.ConsumerSecret)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		return nil, err
	}
	log.Printf("User's ACCOUNT:\n%+v\n", user)
	return client, nil
}

// TwitterConfig ...
func TwitterConfig() *twitter.Stream {
	creds := Credentials{
		AccessToken:       config.TwitterAccessToken,
		AccessTokenSecret: config.TwitterAccessTokenSecret,
		ConsumerKey:       config.TwitterConsumerKey,
		ConsumerSecret:    config.TwitterConsumerSecret,
	}

	log.Printf("Credentials: \n%+v\n", creds)

	client, err := getClient(&creds)
	if err != nil {
		log.Println("Error getting Twitter Client")
		log.Fatal(err)
	}

	log.Printf("%+v\n", client)
	stream, err := client.Streams.Filter(&twitter.StreamFilterParams{
		Follow: []string{config.TwitterID}, // 921550402268606465
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stream:")
	log.Printf("%+v\n", stream)
	return stream
}

// ListenAndNotifyTweets ...
func ListenAndNotifyTweets(bot *tgbotapi.BotAPI, stream *twitter.Stream,
	collection *mongo.Collection) {

	demux := twitter.NewSwitchDemux()
	demux.Tweet = func(tweet *twitter.Tweet) {
		if strings.ToLower(tweet.User.ScreenName) != strings.ToLower(config.TwitterHandle) {
			return
		}
		text := fmt.Sprintf(templates.TwitterMsg, tweet.Text,
			tweet.User.ScreenName, tweet.IDStr)
		users := GetAllChatIDs(collection)
		broadcastTweet(bot, users, text)
	}

	demux.DM = func(dm *twitter.DirectMessage) {
		fmt.Println(dm.SenderID)
	}

	demux.HandleChan(stream.Messages)
}

// GetAllChatIDs ...
func GetAllChatIDs(collection *mongo.Collection) []int64 {
	var usersList []int64
	cur, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Println(err)
		return []int64{}
	}
	for cur.Next(context.TODO()) {
		var user types.User
		if err := cur.Decode(&user); err != nil {
			log.Println(err)
			return []int64{}
		}
		if user.Subscribed {
			usersList = append(usersList, user.ChatID)
		}
	}
	if err := cur.Err(); err != nil {
		log.Println(err)
	}
	cur.Close(context.TODO())
	return usersList
}

func broadcastTweet(bot *tgbotapi.BotAPI, chatIDs []int64, text string) {
	for _, id := range chatIDs {
		msg := tgbotapi.NewMessage(id, text)
		msg.ParseMode = tgbotapi.ModeHTML
		bot.Send(msg)
	}
}
