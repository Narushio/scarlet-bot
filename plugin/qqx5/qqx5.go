package qqx5

import (
	"strings"
)

//var Plugin = bot.NewPlugin("qqx5", "qqx5")

func parseCmdContent(content string) (subCmd string, text string) {
	c := strings.Split(content, " ")
	subCmd = c[0]
	text = c[1]
	return
}
