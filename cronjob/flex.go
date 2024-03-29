package main

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func NewDeluluMessage(year, month, day int, people []string, deluluData map[string]string) *linebot.BubbleContainer {
	choreTaskBox := NewDeluluTaskBox(people, deluluData)
	contents := &linebot.BubbleContainer{
		Type: linebot.FlexContainerTypeBubble,
		Body: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   fmt.Sprintf("%v/%v/%v", year, month, day),
					Weight: "bold",
					Size:   "xxl",
					Color:  "#0000ff",
				},
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   ("暈船打卡🔔"),
					Weight: "bold",
					Size:   "xxl",
					Color:  "#0000ff",
				},
				choreTaskBox,
			},
		},
	}
	return contents
}

func NewDeluluTaskBox(people []string, deluluData map[string]string) *linebot.BoxComponent {
	taskContent := []linebot.FlexComponent{}
	for name, delulu := range deluluData {

		deluluRightComponentContent := []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:      linebot.FlexComponentTypeText,
				Text:      name,
				OffsetTop: "8px",
				Weight:    "bold",
			},
			&linebot.TextComponent{
				Type:      linebot.FlexComponentTypeText,
				Text:      fmt.Sprintf("暈船狀態:%s", delulu),
				OffsetTop: "8px",
				Weight:    "bold",
			},
		}
		if delulu != "沒在暈" {
			deluluRightComponentContent = append(deluluRightComponentContent, &linebot.ButtonComponent{
				Type:      linebot.FlexComponentTypeButton,
				OffsetTop: "10px",
				Style:     "primary",
				Color:     "#00b900",
				Action: &linebot.MessageAction{
					Label: "下船了😎",
					Text:  "下船了😎",
				},
			})
			deluluRightComponentContent = append(deluluRightComponentContent, &linebot.ButtonComponent{
				Type:      linebot.FlexComponentTypeButton,
				Margin:    "2px",
				OffsetTop: "10px",
				Style:     "primary",
				Color:     "#ff0000",
				Action: &linebot.MessageAction{
					Label: "還在暈🤡",
					Text:  "還在暈🤡",
				},
			})
		}
		taskContent = append(taskContent,
			&linebot.BoxComponent{
				Type:   linebot.FlexComponentTypeBox,
				Layout: linebot.FlexBoxLayoutTypeHorizontal,
				Contents: []linebot.FlexComponent{
					&linebot.ImageComponent{
						Type:        linebot.FlexComponentTypeImage,
						URL:         viper.GetString(fmt.Sprintf("PICURL%v", indexOf(name, people)+1)),
						Align:       "start",
						OffsetTop:   "8px",
						AspectRatio: "1:1",
					}, &linebot.BoxComponent{
						// prevent button overflow
						PaddingBottom: "10px",
						Type:          linebot.FlexComponentTypeBox,
						Layout:        linebot.FlexBoxLayoutTypeVertical,
						Contents:      deluluRightComponentContent,
					},
				},
			},
		)
	}

	return &linebot.BoxComponent{
		Type:     linebot.FlexComponentTypeBox,
		Layout:   linebot.FlexBoxLayoutTypeVertical,
		Contents: taskContent,
	}
}

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 //not found.
}
