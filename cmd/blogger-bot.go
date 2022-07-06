package main

import (
	"fmt"
	"github.com/JustinDroege/BloggerBot/pkg/api"
	"github.com/JustinDroege/BloggerBot/pkg/api/message_service"
	"github.com/JustinDroege/BloggerBot/pkg/utils"
	"github.com/bwmarrin/discordgo"
	"github.com/golang-module/carbon/v2"
	"net/http"
	"os"
	"strings"
	"time"
)

var bloggerClient *api.BloggerAPI
var messageBroker *message_service.MessageBroker
var sleepTime int64 = 15
var lastDate = ""

func main() {
	err := setupServices()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = searchForNewPosts()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func setupServices() error {
	apiToken := os.Getenv("BLOGGER_API_TOKEN")
	blogID := os.Getenv("BLOGGER_BLOG_ID")
	discordToken := os.Getenv("DISCORD_BOT_TOKEN")
	discordChannel := os.Getenv("DISCORD_CHANNEL_ID")

	client := &http.Client{}

	temp, err := api.NewBloggerAPI(apiToken, blogID, client, "https://www.googleapis.com/blogger/v3/")

	if err != nil {
		return err
	}

	bloggerClient = temp

	discordSession, err := discordgo.New("Bot " + discordToken)

	if err != nil {
		return err
	}

	discordClient := &message_service.DiscordMessageService{
		DiscordSession:       discordSession,
		ChannelIDs:           &[]string{discordChannel},
		MaximumMessageLength: 2000,
	}

	messageBroker = message_service.New(&[]message_service.MessageService{discordClient})

	return nil
}

func searchForNewPosts() error {
	for {
		time.Sleep(time.Duration(sleepTime) * time.Second)
		posts, err := (*bloggerClient).GetPosts("", strings.Replace(lastDate, "+", "%2B", 1))

		if err != nil {
			return err
		}

		if len(posts.Items) == 0 {
			continue
		}

		if compareDates(lastDate, posts.Items[0].Published) {
			continue
		}

		lastDate = carbon.Parse(posts.Items[0].Published).ToRfc3339String()

		htmlConverted, err := utils.ConvertHtml(posts.Items[0].Content)

		if err != nil {
			return err
		}
		messageBroker.SendMessage(htmlConverted)
	}
}

func compareDates(date1 string, date2 string) bool {
	return carbon.Parse(date1).Gte(carbon.Parse(date2))
}
