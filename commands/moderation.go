package commands

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Not-Cyrus/GoGuardian/utils"
	"github.com/bwmarrin/discordgo"
)

func (cmd *Commands) Ban(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	member, err := s.GuildMember(m.GuildID, m.Mentions[0].ID)
	if err != nil {
		return
	}
	var (
		userRole   = utils.HighestRole(s, m.GuildID, m.Member)
		targetRole = utils.HighestRole(s, m.GuildID, member)
	)

	if !utils.IsAbove(userRole, targetRole) && m.Author.ID != utils.GetGuildOwner(s, m.GuildID) {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("You do not have the proper permissions to ban <@%s>", m.Mentions[0].ID))
		return
	}

	err = s.GuildBanCreateWithReason(m.GuildID, m.Mentions[0].ID, fmt.Sprintf("%s's | Ban CMD Used", s.State.User.Username), 0)

	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**Couldn't ban <@%s>**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Successfully banned <@%s> from the server**", m.Mentions[0].ID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("banned by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Kick(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	member, err := s.GuildMember(m.GuildID, m.Mentions[0].ID)
	if err != nil {
		return
	}

	var (
		userRole   = utils.HighestRole(s, m.GuildID, m.Member)
		targetRole = utils.HighestRole(s, m.GuildID, member)
	)

	if !utils.IsAbove(userRole, targetRole) && m.Author.ID != utils.GetGuildOwner(s, m.GuildID) {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**you do not have the proper permissions to kick <@%s>**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})
		return
	}

	err = s.GuildMemberDeleteWithReason(m.GuildID, m.Mentions[0].ID, fmt.Sprintf("%s's | Ban CMD Used", s.State.User.Username))

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Couldn't kick <@%s>", m.Mentions[0].ID))
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Successfully kicked <@%s> from the server**", m.Mentions[0].ID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Kicked by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Lockdown(s *discordgo.Session, m *discordgo.Message, ctx *Context) {

	err := s.ChannelPermissionSet(m.ChannelID, m.GuildID, "role", 0, discordgo.PermissionSendMessages)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**Couldn't lock this channel. Maybe check perms?**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})

		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Successfully locked <#%s>**", m.ChannelID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Locked by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) SlowMode(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	seconds, err := strconv.Atoi(ctx.Fields[0])
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**You have to specify a number.**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})

		return
	}

	channel, err := s.ChannelEditComplex(m.ChannelID, &discordgo.ChannelEdit{
		RateLimitPerUser: seconds,
	})

	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**Couldn't set slowmode on this channel. Maybe check perms?**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},

			Color: 0x36393F,
		})
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Set <#%s> slowmode to %d seconds**", channel.ID, seconds),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Slowmode set by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) Unban(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	var (
		nukebot    int
		unbancount int
	)

	bans, err := s.GuildBans(m.GuildID)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**Could not fetch the guild bans**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})

		return
	}

	for _, ban := range bans {
		if !strings.Contains(ban.Reason, s.State.User.Username) {
			s.GuildBanDelete(m.GuildID, ban.User.ID)
			unbancount++
			continue
		}
		nukebot++
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Unbanned:**\n*`[ %d ]`* Users\n\n**%s's Bans:**\n*`[ %d ]`* Users", unbancount, s.State.User.Username, nukebot),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("MassUn Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}



func (cmd *Commands) UnLockdown(s *discordgo.Session, m *discordgo.Message, ctx *Context) {

	err := s.ChannelPermissionSet(m.ChannelID, m.GuildID, "role", discordgo.PermissionSendMessages, 0)
	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**Couldn't unlock this channel. Maybe check perms**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Successfully unlocked <#%s>**", m.ChannelID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Unlocked by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	})
}

func (cmd *Commands) UnSlowMode(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	channel, err := s.ChannelEditComplex(m.ChannelID, &discordgo.ChannelEdit{
		RateLimitPerUser: 0,
	})

	if err != nil {
		s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
			Description: fmt.Sprintf("**Couldn't turn off slowmode on this channel. Maybe check perms?**", m.Mentions[0].ID),
			Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
			Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

			Color: 0x36393F,
		})
		return
	}
	s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Description: fmt.Sprintf("**Successfully disabled slowmode for <#%s>**", channel.ID),
		Footer:      &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Slowmode turned off by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),
		Color: 0x36393F,
	})
}
