package main

import (
	"github.com/Noy/DiscordBotGo/utils"
)

func bingo(msg BotMessage) {
	if utils.CaseInsensitiveContains(msg.Message, "bingo") {
		count++
		msg.SendMessagef(`The word "bingo" has been said %d times`, count)
	} else {return}
}