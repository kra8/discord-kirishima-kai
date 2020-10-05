package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

const tokenfile = "token"
const NotifyChannelName = "notify-voice-join"

var voiceActiveUsers map[string]string

func init() {
	voiceActiveUsers = map[string]string{}
}

func main() {
	discord, err := discordgo.New("Bot " + getToken())
	if err != nil {
		panic(err)
	}

	discord.AddHandler(ready)
	discord.AddHandler(voiceStateUpdate)

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("NVJ is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc
	// Cleanly close down the Discord session.
	discord.Close()
	fmt.Println("\nGoodbye.")
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	fmt.Println("now on ready!")
}

func voiceStateUpdate(session *discordgo.Session, voiceState *discordgo.VoiceStateUpdate) {
	// get state updated user
	user, err := session.User(voiceState.UserID)
	if err != nil {
		return
	}

	// get joined voice cahnnel
	joinedChannel, err := session.Channel(voiceState.ChannelID)
	if err != nil {
		delete(voiceActiveUsers, user.ID)
		return
	}

	if channelID, ok := voiceActiveUsers[user.ID]; ok {
		if joinedChannel.ID == channelID {
			return
		}
	}

	// get default text cannel
	channels, err := session.GuildChannels(voiceState.GuildID)
	if err != nil {
		return
	}

	var defaultCannel *discordgo.Channel
	for _, channel := range channels {
		if channel.Name == NotifyChannelName {
			defaultCannel = channel
			break
		}
	}

	if defaultCannel == nil {
		return
	}

	fmt.Println("NOTIFY:" + defaultCannel.Name)

	notifyMessageText := joinedChannel.Name + "に" + user.Username + "さんが参加しました"
	session.ChannelMessageSend(defaultCannel.ID, notifyMessageText)
	voiceActiveUsers[user.ID] = joinedChannel.ID

	fmt.Println(notifyMessageText)
}

func getToken() string {
	text, err := ioutil.ReadFile(tokenfile)
	if err != nil {
		panic("Not found tokenfile.")
	}

	return strings.Trim(string(text), "\n")
}
