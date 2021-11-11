package helpers

import (
	"encoding/json"
	"github.com/anakin0xc06/sentinel_query_bot/config"
	"net/http"

	"github.com/anakin0xc06/sentinel_query_bot/types"
	"gopkg.in/telegram-bot-api.v4"
)

// GetUserName ...
func GetUserName(u tgbotapi.Update) string {
	var username string
	if u.CallbackQuery != nil {
		username = u.CallbackQuery.Message.Chat.UserName
	}
	if u.Message != nil {
		username = u.Message.Chat.UserName
	}
	return username
}

// GetChatID ...
func GetChatID(u tgbotapi.Update) int64 {
	var chatID int64
	if u.CallbackQuery != nil {
		chatID = u.CallbackQuery.Message.Chat.ID
	}
	if u.Message != nil {
		chatID = u.Message.Chat.ID
	}
	return chatID
}

//GetMsgID ...
func GetMsgID(u tgbotapi.Update) int {
	var MsgID int
	if u.CallbackQuery != nil {
		MsgID = u.CallbackQuery.Message.MessageID
	}
	if u.Message != nil {
		MsgID = u.Message.MessageID
	}
	return MsgID
}

// SendMessage ...
func SendMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, text string,
	mode string, btns ...tgbotapi.InlineKeyboardMarkup) {

	if update.Message != nil {
		msg := tgbotapi.NewMessage(GetChatID(update), text)
		if len(btns) > 0 {
			msg.ReplyMarkup = btns[0]
		}
		msg.ParseMode = tgbotapi.ModeMarkdown
		if mode != "" {
			msg.ParseMode = mode
		}
		bot.Send(msg)
		return
	}
	if len(btns) > 0 {
		msg := tgbotapi.NewEditMessageText(GetChatID(update), GetMsgID(update), text)
		msg.ReplyMarkup = &btns[0]
		msg.ParseMode = tgbotapi.ModeMarkdown
		if mode != "" {
			msg.ParseMode = mode
		}
		bot.Send(msg)
		return
	}
	msg := tgbotapi.NewMessage(GetChatID(update), text)
	msg.ParseMode = mode
	bot.Send(msg)
	return
}

// GetDVPNNodesList ...
func GetDVPNNodesList() (types.NodesList, error) {
	var body types.DVPNListResponse
	resp, err := http.Get(config.RestApiURL+ "/nodes")
	if err != nil {
		return types.NodesList{}, err
	}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		return types.NodesList{}, err
	}
	defer resp.Body.Close()
	return body.Result, err
}
