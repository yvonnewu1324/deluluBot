package main

import (
	"deluluBot/helper"
	"fmt"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func main() {
	// Get env
	helper.NewViper()

	// Init line bot
	bot, err := linebot.New(viper.GetString("CHANNELSECRET"), viper.GetString("CHANNELTOKEN"))
	if err != nil {
		helper.LogError("Init Bot Error", err)
	}
	// Compute day of delulu
	year, month, day := time.Now().Date()
	// Send chores notification to group
	var messages []linebot.SendingMessage
	dData := make(map[string]string)
	dData["Helen"] = "沒在暈"
	dData["Yvonne"] = "第Ｎ天"
	contents := NewDeluluMessage(year, int(month), day, []string{"Helen", "Yvonne"}, dData)
	m := linebot.NewFlexMessage(fmt.Sprintf("%v/%v/%v暈船打卡🔔", year, month, day), contents)
	messages = append(messages, m)
	_, err = bot.PushMessage(viper.GetString("GROUPID"), messages...).Do()
	if err != nil {
		helper.LogError("Send Delulu Notify Error", err)
		return
	}
}
