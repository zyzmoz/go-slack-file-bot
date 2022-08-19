package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channellArr := []string{os.Getenv("CHANNEL_ID")}

	fileArr := []string{"file.png"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channellArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)

		if err != nil {
			fmt.Print("%s\n", err)
			return
		}

		fmt.Println("Name: %s, URL: %s", file.Name, file.URL)

	}
}
