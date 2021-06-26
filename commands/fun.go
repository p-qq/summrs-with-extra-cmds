package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type nekoapi struct {
	Url string `json:"url"`
}

func (cmd *Commands) Slap(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/slap")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s *Slapped* %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Poke(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/poke")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s *Poked* %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Tickle(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/tickle")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s *Tickled* %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Kiss(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/kiss")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s *Kissed* %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
func (cmd *Commands) Hug(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/hug")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s Gave %s a hug", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
func (cmd *Commands) Pat(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/pat")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s *Patted* %s *On their head*", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
func (cmd *Commands) Cuddle(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/cuddle")
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
	var send nekoapi
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s *is cuddling with* %s", m.Author.Username, m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
