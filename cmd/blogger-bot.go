package main

import (
	"fmt"
	"github.com/JustinDroege/BloggerBot/pkg/api"
	"github.com/JustinDroege/BloggerBot/pkg/api/message_service"
	"github.com/JustinDroege/BloggerBot/pkg/utils"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
	"time"
)

var bloggerClient *api.BloggerAPI
var messageBroker *message_service.MessageBroker
var sleepTime int64 = 10
var lastPostId string = ""

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
		posts, err := (*bloggerClient).GetPosts("", "")

		if err != nil {
			return err
		}

		if len(posts.Items) == 0 {
			continue
		}

		if lastPostId == posts.Items[0].Id {
			continue
		}

		lastPostId = posts.Items[0].Id

		htmlConverted, err := utils.ConvertHtml(posts.Items[0].Content)

		if err != nil {
			return err
		}

		messageBroker.SendMessage(htmlConverted)
	}
}
