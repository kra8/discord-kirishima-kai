# Kirishima Kai

誰かがボイスチャンネルに参加したとき、テキストチャンネルに通知するDiscordのBotです。 

以前DiscordPHPを利用した同様のBotを作りましたが、ライブラリのサポートが終了したため、リメイクしました。  

# Usage
以下のリンクからあなたのギルドにこのBotを追加できます。  
https://discordapp.com/api/oauth2/authorize?client_id=376220178742116363&scope=bot&permissions=0

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
