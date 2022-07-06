# discord_google_blogger_bot
A discord bot which checks for new posts on specified blogs and post them in specified discord channels

This is my first program written in the language GO in order to learn some stuff about this language. Its a simple Discord bot which checks for new posts on a given Google blog: https://www.blogger.com/ and writes them formated in a discord channel

## Requirements
### Discord bot
In order to use this bot, you need to setup a discord bot: https://discordpy.readthedocs.io/en/stable/discord.html
The discord bot should have a least following permissions: ![permissions](https://i.imgur.com/RUCHlMl.png)

### Go
To run the bot, you need at least GO Version 1.18: https://go.dev/dl/
In the next days, i will push a Docker image to dockerhub. You are free to use the providen dockerfile

## Setup
Just compile the bot with the following command: ```go build ./cmd/blogger-bot.go```. Its should create an executable which you can execute on your current OS. The application needs some environment variables:
| Name of environment variable  | info  |
|---|---|
|  BLOGGER_API_TOKEN | An api token for the google blogger api  |
|  BLOGGER_BLOG_ID | The id from the blog which you want to monitor  |
| DISCORD_BOT_TOKEN  | The discord bot token which you can see, if you created a discord bot like mentioned in the 'Requirements' section  |
| DISCORD_CHANNEL_ID | The discord channel id which the bot can use to send the blogs |


## In the future
In the future, i want to add more features:
- Setup the bot with discord commands (to avoid the env variable DISCORD_CHANNEL_ID)
- Allow multiple guilds and channels
- Use a database to store the configuration (user, guild, channel)
