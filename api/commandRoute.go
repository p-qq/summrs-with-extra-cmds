package api

import (
	"github.com/Not-Cyrus/GoGuardian/commands"
	"github.com/bwmarrin/discordgo"
)

func init() {

	commandRoute.Add("antiinvite", commandRoute.AntiInvite, &commands.Config{
		Alias:        []string{"antiinv", "noinvite"},
		Cooldown:     3,
		OwnerOnly:    true,
		RequiresArgs: true,
	})

	commandRoute.Add("avatar", commandRoute.Avatar, &commands.Config{
		Alias:           []string{"av", "pfp", "icon"},
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("invert", commandRoute.Invert, &commands.Config{
		Alias:           []string{"inve"},
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("glass", commandRoute.Glass, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("gay", commandRoute.Gay, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("waste", commandRoute.Wasted, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("cum", commandRoute.Cum, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("simp", commandRoute.Simp, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("horny", commandRoute.Horny, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("spank", commandRoute.Spank, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("8ball", commandRoute.Eightball, &commands.Config{
		Cooldown:        3,
		
	})
	commandRoute.Add("fact", commandRoute.Fact, &commands.Config{
		Cooldown:        3,
		
	})
	commandRoute.Add("why", commandRoute.Why, &commands.Config{
		Cooldown:        3,
		
	})
	commandRoute.Add("avgen", commandRoute.AvGen, &commands.Config{
		Alias:    []string{"gen", "agen"},
		Cooldown:        3,
		
	})
	commandRoute.Add("wallpaper", commandRoute.Wallpaper, &commands.Config{
		Alias:    []string{"wgen", "wpgen"},
		Cooldown:        3,
		
	})
	commandRoute.Add("bj", commandRoute.BJ, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("anal", commandRoute.Anal, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})

	commandRoute.Add("lesb", commandRoute.Les, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("slap", commandRoute.Slap, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("kiss", commandRoute.Kiss, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("pat", commandRoute.Pat, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("hug", commandRoute.Hug, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("tickle", commandRoute.Tickle, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("poke", commandRoute.Poke, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("cuddle", commandRoute.Cuddle, &commands.Config{
		Cooldown:        3,
		RequiresMention: true,
	})
	commandRoute.Add("ban", commandRoute.Ban, &commands.Config{
		Cooldown:        2,
		RequiresMention: true,
		Perms:           discordgo.PermissionBanMembers,
	})

	commandRoute.Add("banner", commandRoute.ServerBanner, &commands.Config{
		Alias:    []string{"serverbanner", "sbanner"},
		Cooldown: 1,
	})

	commandRoute.Add("botinfo", commandRoute.BotInfo, &commands.Config{
		Cooldown: 4,
	})

	commandRoute.Add("credits", commandRoute.Credits, &commands.Config{
		Cooldown: 1,
	})


	commandRoute.Add("help", commandRoute.Help, &commands.Config{
		Cooldown: 3,
	})

	commandRoute.Add("invite", commandRoute.Invite, &commands.Config{
		Cooldown: 1,
	})

	commandRoute.Add("kick", commandRoute.Kick, &commands.Config{
		Cooldown:        2,
		RequiresMention: true,
		Perms:           discordgo.PermissionBanMembers,
	})

	commandRoute.Add("lockdown", commandRoute.Lockdown, &commands.Config{
		Alias:    []string{"lock"},
		Cooldown: 1,
		Perms:    discordgo.PermissionManageChannels,
	})

	commandRoute.Add("logchannel", commandRoute.LoggingChannel, &commands.Config{
		Alias:     []string{"setlogs", "log"},
		Cooldown:  5,
		OwnerOnly: true,
	})


	commandRoute.Add("massunban", commandRoute.Unban, &commands.Config{
		Alias:    []string{"unbanall","massun","mu"},
		Cooldown: 30,
		Perms:    discordgo.PermissionBanMembers,
	})

	commandRoute.Add("membercount", commandRoute.MemberCount, &commands.Config{
		Alias:    []string{"mc", "members"},
		Cooldown: 1,
	})

	commandRoute.Add("nuke", commandRoute.Nuke, &commands.Config{
		Alias:    []string{"nk"},
		Cooldown: 30,
		Perms:    discordgo.PermissionManageChannels,
	})

	commandRoute.Add("ping", commandRoute.Ping, &commands.Config{
		Alias:    []string{},
		Cooldown: 5,
	})

	commandRoute.Add("prefix", commandRoute.Prefix, &commands.Config{
		Alias:        []string{"setprefix"},
		Cooldown:     3,
		OwnerOnly:    true,
		RequiresArgs: true,
	})

	commandRoute.Add("servericon", commandRoute.ServerIcon, &commands.Config{
		Alias:    []string{"serverpfp", "sicon", "serverpic"},
		Cooldown: 1,
	})

	commandRoute.Add("serverinfo", commandRoute.ServerInfo, &commands.Config{
		Alias:    []string{"si"},
		Cooldown: 1,
	})

	commandRoute.Add("settings", commandRoute.Settings, &commands.Config{
		Cooldown: 1,
	})


	commandRoute.Add("slowmode", commandRoute.SlowMode, &commands.Config{
		Cooldown:     1,
		RequiresArgs: true,
		Perms:        discordgo.PermissionManageChannels,
	})

	commandRoute.Add("unlockdown", commandRoute.UnLockdown, &commands.Config{
		Alias:    []string{"unlock"},
		Cooldown: 1,
		Perms:    discordgo.PermissionManageChannels,
	})

	commandRoute.Add("unslowmode", commandRoute.UnSlowMode, &commands.Config{
		Cooldown: 1,
		Perms:    discordgo.PermissionManageChannels,
	})

	commandRoute.Add("unwhitelist", commandRoute.Unwhitelist, &commands.Config{
		Alias:           []string{"delwl", "removewl", "dewhitelist"},
		Cooldown:        3,
		OwnerOnly:       true,
		RequiresMention: true,
	})

	commandRoute.Add("userinfo", commandRoute.UserInfo, &commands.Config{
		Alias:           []string{"whois"},
		Cooldown:        1,
		RequiresMention: true,
	})

	commandRoute.Add("whitelist", commandRoute.Whitelist, &commands.Config{
		Alias:           []string{"wl", "addwhitelist", "bypass"},
		Cooldown:        3,
		OwnerOnly:       true,
		RequiresMention: true,
	})

	commandRoute.Add("whitelisted", commandRoute.ViewWhitelisted, &commands.Config{
		Alias:           []string{"wled"},
		Cooldown:        3,
		WhitelistedOnly: true,
	})

}
