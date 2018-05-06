package ext

import (
	"gotgbot/types"
	"log"
	"encoding/json"
	"strconv"
	"net/url"
	"github.com/pkg/errors"
)

// TODO: reply version
// TODO: no notif version
func (b Bot) SendGame(chatId int, gameShortName string) (*Message, error) {
	v := url.Values{}
	v.Add("chat_id", strconv.Itoa(chatId))
	v.Add("game_short_name", gameShortName)

	r := Get(b, "sendGame", v)
	if !r.Ok {
		return nil, errors.New(r.Description)
	}

	return b.ParseMessage(r.Result), nil
}


// TODO Check return - message or bool?
// TODO: Force version
// TODO: possible error
func (b Bot) SetGameScore(userId int, score int, chatId int, messageId int) bool {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(userId))
	v.Add("score", strconv.Itoa(score))
	v.Add("chat_id", strconv.Itoa(chatId))
	v.Add("message_id", strconv.Itoa(messageId))

	r := Get(b, "setGameScore", v)
	if !r.Ok {
		log.Fatal("You done goofed, API Res for setGameScore was not OK")
	}

	var bb bool
	json.Unmarshal(r.Result, &bb)

	return bb
}

// TODO Check return - message or bool?
// TODO: Force version
// TODO: Possible error
func (b Bot) SetGameScoreInline(userId int, score int, inlineMessageId string) bool {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(userId))
	v.Add("score", strconv.Itoa(score))
	v.Add("inline_message_id", inlineMessageId)

	r := Get(b, "setGameScore", v)
	if !r.Ok {
		log.Fatal("You done goofed, API Res for setGameScore was not OK")
	}

	var bb bool
	json.Unmarshal(r.Result, &bb)

	return bb
}

func (b Bot) GetGameHighScores(userId int, chatId int, messageId int) []types.GameHighScore {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(userId))
	v.Add("chat_id", strconv.Itoa(chatId))
	v.Add("message_id", strconv.Itoa(messageId))

	r := Get(b, "getGameHighScores", v)
	if !r.Ok {
		log.Fatal("You done goofed, API Res for getGameHighScores was not OK")
	}

	var ghs []types.GameHighScore
	json.Unmarshal(r.Result, &ghs)

	return ghs
}

func (b Bot) GetGameHighScoresInline(userId int, inlineMessageId string) []types.GameHighScore {
	v := url.Values{}
	v.Add("user_id", strconv.Itoa(userId))
	v.Add("inline_message_id", inlineMessageId)

	r := Get(b, "getGameHighScores", v)
	if !r.Ok {
		log.Fatal("You done goofed, API Res for getGameHighScores was not OK")
	}

	var ghs []types.GameHighScore
	json.Unmarshal(r.Result, &ghs)

	return ghs
}