package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Not-Cyrus/GoGuardian/api"
	"github.com/bwmarrin/discordgo"
)

func main() {

	token := "ODM3MTU5MDc1MzI2MDY2NzUw.YIofJQ.7LvVw3wS4RM_hJl1QxIubPjb6YU"

	fmt.Print("Shards: ")
	fmt.Scan(&dshard)

	req, _ := http.NewRequest("GET", "https://discord.com/api/v8/gateway/bot", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bot %s", token))
	res, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Printf("[Sharding Error]: %s\n", err.Error())
		return
	}
	defer res.Body.Close()

	gresponse := &discordgo.GatewayBotResponse{}

	json.NewDecoder(res.Body).Decode(&gresponse)
	if err != nil {
		fmt.Printf("[Decode Error]: %s\n", err.Error())
		return
	}

	var shardCount = gresponse.Shards

	if shardCount < 2 {
		shardCount = dshard
	}

	bot.Sessions = make([]*discordgo.Session, shardCount)

	for s := 0; s < shardCount; s++ {
		bot.Shard(token, shardCount, s)
		bot.Run()
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	bot.Stop()
}

var (
	bot    = api.Bot{}
	dshard int
	token  string
)
