package commands

import (
	"fmt"
	"strings"
	"time"
	"github.com/Not-Cyrus/GoGuardian/database"
	"github.com/bwmarrin/discordgo"
	"go.mongodb.org/mongo-driver/bson"
)

func (cmd *Commands) AntiInvite(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	if !(ctx.Fields[0] == "on" || ctx.Fields[0] == "off") {
		return
	}

	if _, err := database.Database.SetData(m.GuildID, "anti-invite", ctx.Fields[0]); err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Set Anti-Invite to** `%s`", ctx.Fields[0]),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},

		Color: 0x36393F,
	})

}

func (cmd *Commands) LoggingChannel(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	if set, err := database.Database.SetData(m.GuildID, "log-channel", m.ChannelID); !set {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Set Logs to <#%s>**", m.ChannelID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},

		Color: 0x36393F,
	})
}

func (cmd *Commands) Prefix(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	if set, err := database.Database.SetData(m.GuildID, "prefix", ctx.Fields[0]); !set {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("set prefix to: `%s`", ctx.Fields[0]),
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username)},
		Color:  0x36393F,
	})
}

func (cmd *Commands) Settings(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	data, err := database.Database.FindData(m.GuildID)
	guild, _ := s.State.Guild(m.GuildID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	var (
		embed = &discordgo.MessageEmbed{
			Title:  fmt.Sprintf("%s's settings", guild.Name),
			Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username)},
			Color:  0x36393F,
		}
		tempValue string
	)

	for index, value := range data {
		if index == "users" || index == "_id" || index == "guild_id" {
			continue
		}

		switch value.(string) {
		case "on":
			tempValue = "<a:check:837434319714385960>"

		case "off":
			tempValue = "<a:ohnoits:837441036380143646>"

		case "nil":
			tempValue = "<a:ohnoits:837441036380143646>"

		default:
			tempValue = value.(string)
			if index == "log-channel" {
				tempValue = fmt.Sprintf("<#%s>", value.(string))
			}
		}

		embed.Fields = append(embed.Fields, &discordgo.MessageEmbedField{
			Name:   index,
			Value:  tempValue,
			Inline: false,
		})

	}
	s.ChannelMessageSendEmbed(m.ChannelID, embed)
}

func (cmd *Commands) Whitelist(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	if whitelisted, err := database.Database.SetWhitelist(m.GuildID, m.Mentions[0], true); !whitelisted {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Added %s to Whitelist**", m.Mentions[0].Username),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Unwhitelist(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	if whitelisted, err := database.Database.SetWhitelist(m.GuildID, m.Mentions[0], false); !whitelisted {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Removed %s From Whitelist**", m.Mentions[0].Username),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}

func (cmd *Commands) ViewWhitelisted(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	data, err := database.Database.FindData(m.GuildID)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, err.Error())
		return
	}

	var whitelistedUsers []string

	for _, userID := range data["users"].(bson.A) {
		user, _ := s.User(userID.(string))
		whitelistedUsers = append(whitelistedUsers, fmt.Sprintf("<a:supercoolcheck:837407468086755351> | %s#%s", user.Username, user.Discriminator))
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:       "Whitelisted Users",
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Description: strings.Join(whitelistedUsers, "\n"),
		Thumbnail:   &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color:       0x36393F,
	})
}
