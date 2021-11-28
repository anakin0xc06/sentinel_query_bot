package config

import "os"

var (
	RestApiURL    = os.Getenv("REST_API_URL")   // "https://lcd.sentinel.co"
	RpcURL        = os.Getenv("RPC_URL")        // "https://rpc.sentinel.co"
	TwitterHandle = os.Getenv("TWITTER_HANDLE") // "sentinel_co"
	TwitterID     = os.Getenv("TWITTER_ID")     // "921550402268606465"
	MediumHandle  = os.Getenv("MEDIUM_HANDLE")  // "sentinel"
	RedditHandle  = os.Getenv("REDDIT_HANDLE")  // "r/SENT"
	CoingeckoID   = os.Getenv("COINGECKO_ID")   // sentinel

	// coin details
	CoinDecimals     = 6
	CoinDenom        = "udvpn"
	CoinDisplayDenom = "DVPN"
	AddressPrefix    = "sent"
)

// Twitter Config

var (
	TwitterAccessToken       = os.Getenv("TWITTER_ACCESS_TOKEN")
	TwitterAccessTokenSecret = os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	TwitterConsumerKey       = os.Getenv("TWITTER_CONSUMER_API_KEY")
	TwitterConsumerSecret    = os.Getenv("TWITTER_CONSUMER_API_SECRET_KEY")
)

func init() {
	if len(TwitterAccessToken) < 1 ||
		len(TwitterAccessTokenSecret) < 1 ||
		len(TwitterConsumerKey) < 1 ||
		len(TwitterConsumerSecret) < 1 {

		panic("Please configure twitter credentials")
	}
}
