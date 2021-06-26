package commands

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/bwmarrin/discordgo"
)

type api struct {
	Url     string `json:"url"`
	MSG     string `json:"response"`
	Fact    string `json:"fact"`
	Why     string `json:"why"`
}

func (cmd *Commands) Eightball(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/8ball")
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
	var send api
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s's Response:\n", s.State.User.Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Description: fmt.Sprintf(send.MSG),
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Wallpaper(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/wallpaper")
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
	var send api
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s generated Wallpaper For **%s**", s.State.User.Username, m.Author.Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
func (cmd *Commands) AvGen(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/img/avatar")
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
	var send api
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s generated a pfp for **%s**", s.State.User.Username, m.Author.Username),
		Image:  &discordgo.MessageEmbedImage{URL: send.Url},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
//https://nekos.life/api/v2/fact

func (cmd *Commands) Fact(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/fact")
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
	var send api
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s's fact OTD:", s.State.User.Username),
		Description: fmt.Sprintf("**True Fact**:\n\n%s",send.Fact),
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
//https://nekos.life/api/v2/why
func (cmd *Commands) Why(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	url := fmt.Sprintf("https://nekos.life/api/v2/why")
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
	var send api
	if err := json.NewDecoder(resp.Body).Decode(&send); err != nil {

	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s's Daily Life question:", s.State.User.Username),
		Description: fmt.Sprintf("**Can You Answer This?**:\n\n%s",send.Why),
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Wasted(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("Wasted: **%s**", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/wasted/?avatar=%s", m.Mentions[0].AvatarURL("2048"))},
		
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
func (cmd *Commands) Trigger(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("Triggered: **%s**", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/triggered?avatar=%s", m.Mentions[0].AvatarURL("2048"))},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}

func (cmd *Commands) Horny(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("**%s's** Horny License", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/horny?avatar=%s", m.Mentions[0].AvatarURL("2048"))},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}
func (cmd *Commands) Simp(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("**%s's** Simp Card", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/simpcard?avatar=%s", m.Mentions[0].AvatarURL("2048"))},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}
func (cmd *Commands) Invert(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("inverted: **%s**", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/invert?avatar=%s", m.Mentions[0].AvatarURL("2048"))},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}

func (cmd *Commands) Gay(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("Turned **%s** gay", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/gay?avatar=%s", m.Mentions[0].AvatarURL("2048"))},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}
func (cmd *Commands) Glass(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("Glassed: %s", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: fmt.Sprintf("https://some-random-api.ml/canvas/glass?avatar=%s",m.Mentions[0].AvatarURL("2048"))},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}

//https://some-random-api.ml/canvas/glass
//https://some-random-api.ml/canvas/simpcard?avatar=

//https://some-random-api.ml/