# Kirishima Kai

誰かがボイスチャンネルに参加したとき、テキストチャンネルに通知するDiscordのBotです。 

以前DiscordPHPを利用した同様のBotを作りましたが、ライブラリのサポートが終了したため、リメイクしました。  

# Usage
以下のリンクからあなたのギルドにこのBotを追加できます。  
https://discordapp.com/api/oauth2/authorize?client_id=440799039835996170&permissions=0&redirect_uri=https%3A%2F%2Fgithub.com%2Fkra8%2Fdiscord-kirishima-kai&scope=bot

もしくは、このBotを自身のアプリケーションとして動かすことができます。
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
