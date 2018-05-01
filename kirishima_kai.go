package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

const tokenfile = "token"

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
	discord.AddHandler(messageCreate)
	discord.AddHandler(voiceStateUpdate)

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Kirishima is now running.  Press CTRL-C to exit.")
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

func messageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == session.State.User.ID {
		return
	}

	for _, user := range message.Mentions {
		if user.ID == session.State.User.ID {
			session.ChannelMessageSend(message.ChannelID, getSerif())
		}
	}
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
	cannels, err := session.GuildChannels(voiceState.GuildID)
	if err != nil {
		return
	}
	defaultCannel := cannels[0]

	session.ChannelMessageSend(defaultCannel.ID, joinedChannel.Name+"に"+user.Username+"さんが参加しました")
	voiceActiveUsers[user.ID] = joinedChannel.ID
}

func getToken() string {
	text, err := ioutil.ReadFile(tokenfile)
	if err != nil {
		panic("Not found tokenfile.")
	}

	return strings.Trim(string(text), "\n")
}

func getSerif() string {
	serifs := [...]string{
		"さっ、早くご命令を。司令？",
		"ご命令を、司令",
		"そのー、何度もつつかれるのは、何でしょう。新たなコマンドなんでしょうか？",
		"備えあれば憂いなし、です。",
	}

	rand.Seed(time.Now().UnixNano())
	return serifs[rand.Intn(len(serifs))]
}
