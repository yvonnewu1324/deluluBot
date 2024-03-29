package helper

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func NewHttpHandler(bot *linebot.Client) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			LogError("Bot Parse Error", err)
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		// random reply message to keyword "還在暈🤡" "下船了😎"
		for _, event := range events {
			if event.Type == linebot.EventTypeJoin {
				LogInfo(fmt.Sprintf("bot joined this group! groupID:%s", event.Source.GroupID))
			}
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					replyDelulu := [6]string{"就你在丑 🤡", "寶打咩 👹", "快點下船 🫠", "別太愛 😍", "來 解暈藥 💊", "你超愛 😂"}
					replyOver := [5]string{"恭喜下船 🫡", "喝一杯 🍻", "good job 🙌", "yes queen 🫶", "nice 🤙"}
					if message.Text == "還在暈🤡" {
						profile, err := bot.GetProfile(event.Source.UserID).Do()
						if err != nil {
							LogError("Get Profile Error", err)
							return
						}
						replyMessage := profile.DisplayName + " " + replyDelulu[rand.Intn(len(replyDelulu))]
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							LogError("Send Reply Done Messege Error", err)
						}
					}
					if message.Text == "下船了😎" {
						profile, err := bot.GetProfile(event.Source.UserID).Do()
						if err != nil {
							LogError("Get Profile Error", err)
							return
						}
						replyMessage := profile.DisplayName + " " + replyOver[rand.Intn(len(replyOver))]
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(replyMessage)).Do(); err != nil {
							LogError("Send Reply Done Messege Error", err)
						}
					}
				case *linebot.StickerMessage:
					// replyMessage := fmt.Sprintf(
					// 	"sticker id is %s, stickerResourceType is %s", message.StickerID, message.StickerResourceType)
					// Seed the random number generator with the current time
					source := rand.NewSource(time.Now().UnixNano())
					randGenerator := rand.New(source)
					// Generate a random number between 52114110 and 52114149
					randomNumber := randGenerator.Intn(52114149-52114110+1) + 52114110
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("11539", fmt.Sprintf("%v", randomNumber))).Do(); err != nil {
						LogError("Send Reply Sticker Messege Error", err)
					}
				}
			}
		}
	}
}
