package commands

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Not-Cyrus/GoGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func (cmd *Commands) Avatar(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Title:  fmt.Sprintf("%s's avatar", m.Mentions[0].Username),
		Image:  &discordgo.MessageEmbedImage{URL: m.Mentions[0].AvatarURL("2048")},
		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) UserInfo(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	member, err := s.GuildMember(m.GuildID, m.Mentions[0].ID)
	if err != nil {
		return
	}

	var (
		memberCreatedAt, _ = discordgo.SnowflakeTimestamp(m.Mentions[0].ID)
		memberJoinedAt, _  = member.JoinedAt.Parse()
		role               = utils.HighestRole(s, m.GuildID, member)
		roleID             = "@everyone"
	)

	if role != nil {
		roleID = fmt.Sprintf("<@&%s>", role.ID)
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Author:    &discordgo.MessageEmbedAuthor{Name: fmt.Sprintf("%s#%s's info:\n", m.Mentions[0].Username, m.Mentions[0].Discriminator)},
		Thumbnail: &discordgo.MessageEmbedThumbnail{URL: m.Mentions[0].AvatarURL("")},

		Fields: []*discordgo.MessageEmbedField{
			{Name: "Username:", Value: m.Mentions[0].Username, Inline: true},
			{Name: "Discriminator:", Value: m.Mentions[0].Discriminator, Inline: true},
			{Name: "ID:", Value: m.Mentions[0].ID, Inline: true},
			{Name: "Account Made On:", Value: memberCreatedAt.Format("01/02/2006"), Inline: true},
			{Name: "Account Joined On:", Value: memberJoinedAt.Format("01/02/2006"), Inline: true},
			{Name: "Bot?", Value: strconv.FormatBool(m.Mentions[0].Bot), Inline: true},
			{Name: "Highest Role:", Value: roleID, Inline: true},
		},

		Footer: &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}
