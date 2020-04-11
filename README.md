# Kirishima Kai

誰かがボイスチャンネルに参加したとき、テキストチャンネルに通知するDiscordのBotです。 

# Usage
- 以下のリンクからあなたのギルドにこのBotを追加
https://discordapp.com/api/oauth2/authorize?client_id=440799039835996170&permissions=0&redirect_uri=https%3A%2F%2Fgithub.com%2Fkra8%2Fdiscord-kirishima-kai&scope=bot

- 通知用のテキストチャンネルを作成します
`notify-voice-join` という名前のテキストチャンネルを作成してください。

### また、このBotを自身のアプリケーションとして動かすことができます。
```
$ git clone https://github.com/kra8/discord-kirishima-kai.git
$ cd discord-kirishima-kai
$ go get github.com/bwmarrin/discordgo
$ make build
$ cp token.example token
> tokenにあなたのアプリケーションのBot tokenを記述してください
$ ./bin/kirishima-kai
```

# LICENCE
[MIT](https://github.com/kra8/discord-kirishima-kai/blob/master/LICENCE)
