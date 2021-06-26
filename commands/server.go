package commands

import (
	"fmt"
	"time"
	"github.com/bwmarrin/discordgo"
)

func (cmd *Commands) MemberCount(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: guild.Name,
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Members", Value: fmt.Sprint(guild.MemberCount), Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},

		Color: 0x36393F,
	})

}

func (cmd *Commands) Nuke(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		return
	}

	_, err = s.ChannelDelete(channel.ID)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "I can't delete that channel.")
		return
	}

	channel, err = s.GuildChannelCreateComplex(m.GuildID, discordgo.GuildChannelCreateData{
		Name:                 channel.Name,
		Type:                 channel.Type,
		Topic:                channel.Topic,
		RateLimitPerUser:     channel.RateLimitPerUser,
		Position:             channel.Position,
		PermissionOverwrites: channel.PermissionOverwrites,
		ParentID:             channel.ParentID,
		NSFW:                 channel.NSFW,
	})
	s.ChannelMessageSendEmbed(channel.ID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s purged channel", m.Author.Username),

		Image: &discordgo.MessageEmbedImage{URL: "https://media.giphy.com/media/oe33xf3B50fsc/giphy.gif"},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
	time.Sleep(2 * time.Second)

}

func (cmd *Commands) ServerBanner(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return
	}

	if len(guild.Banner) == 0 {
		s.ChannelMessageSend(m.ChannelID, "There is no guild banner.")
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s's server banner", guild.Name),
		Image: &discordgo.MessageEmbedImage{URL: discordgo.EndpointGuildBanner(guild.ID, guild.Banner)},

		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username)},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:  0x36393F,
	})

}

func (cmd *Commands) ServerIcon(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return
	}

	if len(guild.IconURL()) == 0 {
		s.ChannelMessageSend(m.ChannelID, "There is no guild icon.")
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s's server icon", guild.Name),
		Image: &discordgo.MessageEmbedImage{
			URL: guild.IconURL(),
		},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})

}

func (cmd *Commands) ServerInfo(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return
	}

	guildTime, _ := discordgo.SnowflakeTimestamp(guild.ID)

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title: fmt.Sprintf("%s's Info", guild.Name),
		Fields: []*discordgo.MessageEmbedField{
			{Name: "Members:", Value: fmt.Sprint(guild.MemberCount), Inline: true},
			{Name: "Roles:", Value: fmt.Sprint(len(guild.Roles)), Inline: true},
			{Name: "Channels:", Value: fmt.Sprint(len(guild.Channels)), Inline: true},
			{Name: "Server Created on", Value: guildTime.Format("01/02/2006"), Inline: true},
			{Name: "Server Owner", Value: fmt.Sprintf("<@%s>", guild.OwnerID), Inline: true},
			{Name: "Server Region", Value: guild.Region, Inline: true},
		},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: guild.IconURL()},
		Color:     0x36393F,
	})

}
