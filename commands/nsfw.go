package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/bwmarrin/discordgo"
)

type nekoapi2 struct {
	Url string `json:"url"`
}

func (cmd *Commands) Anal(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/anal")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var send nekoapi2
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s  *is doing anal wtih*  %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}

func (cmd *Commands) Cum(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/cum")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var send nekoapi2
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s  *nutted inside*  %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})	
}

func (cmd *Commands) BJ(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/bj")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var send nekoapi2
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s  *just gave*  %s  *a blowjob*", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Spank(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/spank")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var send nekoapi2
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s  *spanked*  %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}

func (cmd *Commands) Les(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/les")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}
	defer resp.Body.Close()
	var send nekoapi2
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s  *is showing*  %s  *some sis luv*", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}
