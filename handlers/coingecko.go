package handlers

import (
	"fmt"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"github.com/anakin0xc06/sentinel_query_bot/helpers"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"github.com/anakin0xc06/sentinel_query_bot/types"
	coingecko "github.com/superoo7/go-gecko/v3"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"net/http"
	"strings"
	"time"
)

var cg *coingecko.Client

func init() {
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	cg = coingecko.NewClient(httpClient)
}

func GetCoinGeckoPrice(id string) (*types.Price, error) {
	ids := []string{id}
	vc := []string{"usd"}
	price := &types.Price{}

	coin, err := cg.CoinsID(id, false, false, true, true, false, false)
	if err != nil {
		return nil, err
	}
	sp, err := cg.SimplePrice(ids, vc)
	if err != nil {
		return nil, err
	}
	cgPrice := (*sp)[id]
	price.Symbol = coin.Symbol
	price.Current = cgPrice["usd"]
	price.ATH = coin.MarketData.ATH["usd"]
	price.ATL = coin.MarketData.ATL["usd"]
	price.Low24h = coin.MarketData.Low24["usd"]
	price.High24h = coin.MarketData.High24["usd"]
	price.ChangePercentage24h = coin.MarketData.PriceChange24hInCurrency["usd"]

	fmt.Printf("Price: %v", price)
	return price, nil

}
func GetCoinPriceData(id string) (string, error) {
	p, err := GetCoinGeckoPrice(id)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%.2f", p.Current), err
}

// PriceHandler ...
func PriceHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	p, err := GetCoinGeckoPrice(config.CoingeckoID)
	if err != nil {
		helpers.SendReplyMessage(bot, update, templates.UnableToGetPrice, tgbotapi.ModeHTML)
	}
	helpers.SendReplyMessage(bot, update, fmt.Sprint(p.Current), tgbotapi.ModeHTML)
}

// PriceHandler ...
func PriceStatsHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	p, err := GetCoinGeckoPrice(config.CoingeckoID)
	if err != nil {
		helpers.SendReplyMessage(bot, update, templates.UnableToGetPrice, tgbotapi.ModeHTML)
	}
	if p != nil {
		text := fmt.Sprintf(
			templates.PriceStats,
			config.CoingeckoID,
			strings.ToUpper(p.Symbol),
			p.Current,
			p.High24h,
			p.Low24h,
			p.ChangePercentage24h*100,
			p.ATH,
			p.ATL,
		)
		helpers.SendReplyMessage(bot, update, text, tgbotapi.ModeHTML)
		return
	}
}
