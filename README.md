# BanLogger

BanLogger is a Discord bot which allows game server administrators to easily log and form a record of the offenses of users in Steam games.

## Setup
1. Install Go 1.14 or newer
2. Clone this repository
3. Build the project using `go build cmd/banlogger/main.go`. This will output a binary called `main`. You should rename this to BanLogger or something else which makes sense to you.
4. Create a `.env` file alongside the executable and define the following environment variables:
`DISCORD_BOT_TOKEN` `STEAM_API_KEY` `CHANNEL_ID`

DISCORD_BOT_TOKEN is where you put your Discord bot's token. You can create an application and then a bot at https://discord.com/developers/applications.
For more information, give this page a read: https://www.writebots.com/discord-bot-token/

STEAM_API_KEY is where you put your Steam API Key. You can get your API key at https://steamcommunity.com/dev/apikey.

CHANNEL_ID where you put the ID of the channel you wish the bot to run in. For more info on how to get a Discord channel's ID, check out this guide: https://support.discord.com/hc/en-us/articles/206346498-Where-can-I-find-my-User-Server-Message-ID-

Once all this is done, you should be able to run the bot. Run it like you would any other executable and invite the bot to your server!
