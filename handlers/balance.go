package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"github.com/anakin0xc06/sentinel_query_bot/helpers"
	"github.com/anakin0xc06/sentinel_query_bot/templates"
	"github.com/anakin0xc06/sentinel_query_bot/utils"
	tgbotapi "gopkg.in/telegram-bot-api.v4"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Balances struct {
	Balances []struct {
		Denom  string `json:"denom"`
		Amount string `json:"amount"`
	} `json:"balances"`
}

var balancesURL = fmt.Sprintf("%s/cosmos/bank/v1beta1/balances", config.RestApiURL)

func GetAccountBalances(address string) (*Balances, error) {
	balances := &Balances{}

	url := fmt.Sprintf("%s/%s", balancesURL, address)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return balances, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&balances); err != nil {
		return balances, err
	}
	defer resp.Body.Close()
	fmt.Println(balances)
	return balances, nil
}

func BalanceHandler(bot *tgbotapi.BotAPI, update tgbotapi.Update, address string) {
	if address == "" {
		args := strings.Split(update.Message.CommandArguments(), " ")
		address = args[0]
	}
	if !isValidaAddress(address) {
		text := fmt.Sprintf("invalid address, address should be a valid bech32 address with prefix \"%s\"", config.AddressPrefix)
		helpers.SendReplyMessage(bot, update, text, tgbotapi.ModeHTML)
		return
	}

	fmt.Println("fetching balance for address ", address)
	balances, err := GetAccountBalances(address)
	if err != nil {
		helpers.SendReplyMessage(bot, update, templates.UnableToGetPrice, tgbotapi.ModeHTML)
	}
	text := fmt.Sprintf("\n<b>Address</b>: %s\n", address)
	text += "<b>Balances:</b>\n"
	for _, balance := range balances.Balances {
		b := fmt.Sprintf("<i>%s %s</i>\n", balance.Amount, balance.Denom)
		if balance.Denom == config.CoinDenom {
			amount, err := strconv.ParseFloat(balance.Amount, 64)
			if err != nil {
				helpers.SendReplyMessage(bot, update, templates.UnalbleToGetBalance, tgbotapi.ModeHTML)
				return
			}
			b = fmt.Sprintf(templates.BalanceFormat, amount/math.Pow10(config.CoinDecimals), config.CoinDisplayDenom)
		}
		text += b
	}
	if len(balances.Balances) == 0 {
		text += "\nNo Available Balances\n"
	}
	helpers.SendReplyMessage(bot, update, text, tgbotapi.ModeHTML)
}

func isValidaAddress(addr string) bool {
	prefix, err := utils.DecodeBech32Address(addr)
	if err != nil {
		return false
	}
	if prefix != config.AddressPrefix {
		return false
	}
	return true
}
