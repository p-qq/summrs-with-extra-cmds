package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)



func (cmd *Commands) Help(s *discordgo.Session, m *discordgo.Message, ctx *Context) {
	if len(ctx.Fields) == 0 {
		
		defaultHelp.Description = fmt.Sprintf("**`%shelp [category]`. cmd count: [`42`]** ***Unfinished***\n\n",ctx.Prefix)
		defaultHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		defaultHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
		

		s.ChannelMessageSendEmbed(m.ChannelID, defaultHelp)
		return
	}
	
	defaultHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
	certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}

	switch ctx.Fields[0] {

	case "information":
		certainHelp.Title = "<a:hoodwumpus_nitro:837349614600454154> · Information · <a:hoodwumpus_nitro:837349614600454154>"
		certainHelp.Description = fmt.Sprintf("**%sserverinfo:**\n*`Returns information about the current server`*\n\n**%sbotinfo:**\n*`Shows information about the bot`*\n\n**%suserinfo [@user]:**\n*`Shows informati on about the mentioned user`*\n\n**%savatar [@user]:**\n*`Returns the mentioned users avatar`*\n\n**%smembercount:**\n*`Returns the server's member count`*\n\n**%sbanner:**\n*`Returns the server banner`*\n\n**%sservericon:**\n*`Returns the server icon`*", ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix) //yeah ummm we don't talk about this..
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	case "anti":
		certainHelp.Title = "<a:DiscordPartner:837363864534253700> · Anti Nuke · <a:DiscordPartner:837363864534253700> `Not Done` "
		certainHelp.Description = fmt.Sprintf("**%swhitelist [@user]:**\n*`Whitelists the mentioned user`*\n\n**%sunwhitelist [@user]:**\n*`Dewhitelists the mentioned user`*\n\n**%swhitelisted:**\n*`Shows whitelisted list`*", ctx.Prefix, ctx.Prefix, ctx.Prefix)
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	case "moderation":
		certainHelp.Title = "<a:blob_discord:837373791491194920> · Moderation · <a:blob_discord:837373791491194920>"
		certainHelp.Description = fmt.Sprintf("**%sban [@user]:**\n*`Bans the mentioned user`*\n\n**%skick [@user]:**\n*`Kicks the mentioned user`*\n\n**%spurge [amount]:**\n*`Purges entered amount of messages`*\n\n**%slock:**\n*`Locks the channel`*\n\n**%sunlock:**\n*`Unlocks the channel`*\n\n**%sslowmode [time]:**\n*`Sets the channel slowmode to that time`*\n\n**%sunslowmode:**\n*`Disables slow mode`*\n\n**%smassunban:**\n*`Unbans all members in the server`*", ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix)
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	case "misc":
		certainHelp.Title = "<a:blob_discord:837373791491194920> · Misc · <a:blob_discord:837373791491194920>"
		certainHelp.Description = fmt.Sprintf("**%s8ball:**\n*`generates a 8ball answer`*\n\n**%savgen:**\n*`generates random anime pfp`*\n\n**%swpgen:**\n*`generates random anime wallpaper`*\n\n**%swhy:**\n*`asks you a random question`*\n\n**%sfact:**\n*`drops a random fact for you :)`*\n\n**%ssimp:**\n*`turns user into a simp`*\n\n**%shorny:**\n*`makes the user horny`*\n\n**%sgay:**\n*`turns user gay`*\n\n**%sglass:**\n*`turns user into a simp`*\n\n**%swaste:**\n*`wasted effect on user`*",ctx.Prefix,ctx.Prefix,ctx.Prefix, ctx.Prefix, ctx.Prefix,ctx.Prefix,ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix)
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	case "settings":
		certainHelp.Title = "<a:hoodwumpus_nitro:837349614600454154> · Settings · <a:hoodwumpus_nitro:837349614600454154>"
		certainHelp.Description = fmt.Sprintf("**%sprefix [prefix]:**\n*`Sets the bot prefix`*\n\n**%slogchannel:**\n*`Sets the log channel for all notifications relating to the anti-nuke.`*\n\n**%santiinvite [on/off]:**\n*`Enables/Disables the anti invite system`*", ctx.Prefix, ctx.Prefix, ctx.Prefix)
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	case "fun":
		certainHelp.Title = "<a:DiscordPartner:837363864534253700> · Fun · <a:DiscordPartner:837363864534253700>"
		certainHelp.Description = fmt.Sprintf("**%skiss:**\n*`kisses the user you mention`*\n\n**%shug:**\n*`hugs the user you mention`*\n\n**%sslap:**\n*`slaps the user you mention`*\n\n**%spoke:**\n*`pokes the user you mention`*\n\n**%scuddle:**\n*`cuddles with the user you mention`*\n\n**%stickle:**\n*`tickles the user you mention`*", ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix)
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	case "nsfw":
		certainHelp.Title = "<a:infinite:837411722193010729> · Nsfw · <a:infinite:837411722193010729>"
		certainHelp.Description = fmt.Sprintf("**%sanal:**\n*`anal with user you mention`*\n\n**%scum:**\n*`cums on the user you mention`*\n\n**%sbj:**\n*`gives a blowjob to the user you mention`*\n\n**%sspank**\n*`spanks the user you mention`*\n\n**%slesb**\n*`sex but 4 females only`*", ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix, ctx.Prefix)
		certainHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		certainHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}
	default:
		defaultHelp.Title = fmt.Sprintf(s.State.User.Username)
		defaultHelp.Description = fmt.Sprintf("type `%shelp [category]`. cmds: [`45`] ***Unfinished***", ctx.Prefix)
		defaultHelp.Footer = &discordgo.MessageEmbedFooter{Text: fmt.Sprintf("Requested by: %s", m.Author.Username), IconURL: m.Author.AvatarURL("")}
		defaultHelp.Thumbnail = &discordgo.MessageEmbedThumbnail{URL: s.State.User.AvatarURL("")}

		s.ChannelMessageSendEmbed(m.ChannelID, defaultHelp)
		return
	}

	s.ChannelMessageSendEmbed(m.ChannelID, certainHelp)
}



var (
	certainHelp = &discordgo.MessageEmbed{
		Color: 0x36393F,
	}
	
	defaultHelp = &discordgo.MessageEmbed{
		
		Fields: []*discordgo.MessageEmbedField{
			{Name: "**Information** | [`7`]", Value: "\n`help information`",Inline: true},
			{Name: "**Anti** | [`3`]", Value: "\n`help anti`",Inline: true},
			{Name: "**Moderation** | [`10`]", Value: "\n`help moderation`",Inline: true},
			{Name: "‎‎‎‎‎‏‏‎**Settings** | [`3`]", Value: "\n`help settings`",Inline: true},
			{Name: "**Fun** | [`7`]", Value: "‎‎‎‎‎‏‏‎`help fun`",Inline: true},
			{Name: "**Nsfw** | [`5`]", Value: "\n‏‏‎`help nsfw`",Inline: true},
			{Name: "‎‎‎‎‎**Misc** | [`5`]", Value: "\n‏‏‎`help misc`",Inline: true},
		},
		Timestamp: time.Now().UTC().Format("2006-01-02T15:04:05-0700"),

		Color: 0x36393F,
	}
)
