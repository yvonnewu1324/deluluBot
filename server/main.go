package main

import (
	"deluluBot/helper"
	"fmt"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func main() {
	// Get env
	helper.NewViper()

	// Init line bot with http client
	client := &http.Client{}
	// Do delulu CronJob
	http.HandleFunc("/delulu-cronjob", helper.NewCronJobHandler())
	bot, err := linebot.New(viper.GetString("CHANNELSECRET"), viper.GetString("CHANNELTOKEN"), linebot.WithHTTPClient(client))
	if err != nil {
		fmt.Println("Init Bot Error:", err)
	}
	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", helper.NewHttpHandler(bot))
	// Set server port
	port := viper.GetInt("PORT")
	fmt.Printf("Server is running on :%d\n", port)

	// Start server
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}

}
