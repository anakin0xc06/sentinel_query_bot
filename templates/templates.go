package templates

var (
	// home templates

	WelcomeGreetMsg = `Hello @%s,

*Welcome to the Sentinel Network Query Bot*`
	SelectHelpMsg = "Please select /help to know commands available for this bot"

	SubscribeMsg   = "Successfully subscribed to @%s.\n\nNow you will receive updates from Sentinel Network"
	UnSubscribeMsg = "Successfully unsubscribed to @%s.\n\nNow you will not receive any updates from Sentinel Network"

	// comunity register
	ChannelRegisteredMsg   = "This channel is successfully registered with @%s"
	ChannelUnRegisteredMsg = "This channel is successfully unregistered with @%s"

	// About Sentinel template

	AboutSentinel = `<b>Sentinel Network </b>

	 Website: https://sentinel.co
	 Github: https://github.com/sentinel-official/
	 Twitter: http://twitter.com/sentinel_co`

	// About Bot
	AboutBot = `<b>Sentinel Query Bot</b>

	This bot used to query sentinel network 
     - query price of sentinel dvpn
     - query balances of sentinel account

    comming soon:
     - delegations of address 
     - votes/proposals of address
     - validator updates to saved accounts
     
    developed by <b>Daksha Validator Team</b>.`

	// Help template

	HelpMsg = `
	<b>Here are the available commands and their utility</b>

	1. /subscribe - subscribe to sentinel network social media updates
	2. /unsubscribe - unsubscribe updates
	3. /sentinel - learn about sentinel network
	4. /p - current price of sentinel dvpn from coingecko 
	5. /pricestats - price statistics of sentinel dvpn from coingecko
	6. /balances [addr] - get balances of address 
	7. /about - about bot`

	// social media updates templates

	TwitterMsg = `<b>New Tweet from </b>@sentinel_co

	%s

	link: https://twitter.com/%s/status/%s`

	RetweetMsg   = "<b>New Retweet from %s</b>\n\n%s\nhttps://twitter.com/%s/status/%s"
	MediumPost   = "<b>New Medium Publication from @%s</b>\n\n<b>%s</b>\n\nLink: %s"
	RedditPost   = "<b>New Reddit post from %s</b>\n\n%s\n\nlink: %s"
	TweetFormat  = "<b>@Sentinel_co tweeted on %s</b>\n\n%s"
	RedditFormat = "<b>Post on %s</b>\n\n%s\n\nLink: %s"
	MediumFormat = "<b>Recent Medium Publication from @sentinel_co</b>\n\n<b>%s</b>\n\nLink: %s"

	// coingecko price templates
	PriceStats = `<b>%s</b>

Symbol: $%s
Current Price (USD): %.3f

24h High: %.3f
24h Low: %.3f
24h Change Percentage: %.2f %% 

ATH: %.3f
ATL: %.3f`

	UnableToGetPrice = "unable to fetch price details at the moment, please try again later."

	// account templates
	UnalbleToGetBalance = "unable to get balance at the moment, please try again later."
	BalanceFormat       = "<i>%.2f %s</i>\n"
)
