package commands

import (
	"fmt"
	"time"

	"github.com/Not-Cyrus/GoGuardian/events"

	"github.com/bwmarrin/discordgo"
)

func (cmd *Commands) BotInfo(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: "Bot Info",

		Fields: []*discordgo.MessageEmbedField{
			{Name: "Name:", Value: s.State.User.Username, Inline: false},
			{Name: "Server Count:", Value: fmt.Sprint(events.GuildCount), Inline: false},
			{Name: "User Count:", Value: fmt.Sprint(events.MemberCount), Inline: false},
			{Name: "Ping:", Value: fmt.Sprintf("%s", s.HeartbeatLatency().Round(1*time.Millisecond)), Inline: false},
			{Name: "Golang Version", Value: "v0.22.0", Inline: false},
			{Name: "Invite", Value: fmt.Sprintf("[Click Here](https://discord.com/api/oauth2/authorize?client_id=%s&permissions=8&scope=bot) also join discord.gg/coding for bot support :)", s.State.User.ID), Inline: true},
		},

		Footer:    &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:     0x36393F,
	})
}

func (cmd *Commands) Credits(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Credits",
		Description: "If you would like to come support us in the server then feel free to join click [Invite Me](https://discord.gg/coding) to get invited if discord.gg/coding doesn't work. This bot is unfinished so there will be a lot of bugs here and there.",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Creators:", Value: "Original Src: [!fishgang Cy](https://github.com/Not-Cyrus)\nDevelopers: [pix#0070](https://github.com/eozri)\n[pix#0080](https://github.com/eozri)\nOur Server: [github](https://discord.gg/coding)"},
		},
		Footer:    &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:     0x36393F,
	})
}

func (cmd *Commands) Invite(s *discordgo.Session, m *discordgo.Message, ctx *Context) {

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s's Bots Invite", s.State.User.Username),
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Click Below:", Value: fmt.Sprintf("\n[Click Here](https://discord.com/api/oauth2/authorize?client_id=%s&permissions=8&scope=bot) also join discord.gg/coding for bot support :)", s.State.User.ID), Inline: true},
		},
		Footer:    &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:     0x36393F,
	})
}

func (cmd *Commands) Ping(s *discordgo.Session, m *discordgo.Message, ctx *Context) {

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("%s's ping", s.State.User.Username),
		Description: fmt.Sprintf("Bot Ping: `%s`\n\nCurrent shards `%d`", s.HeartbeatLatency().Round(1*time.Millisecond), s.ShardID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Timestamp:   time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:       0x36393F,
	})
}
