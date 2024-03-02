package discord

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/kehiy/RoboPac/engine"
	"github.com/kehiy/RoboPac/log"
)

func (db *DiscordBot) respondErrMsg(cmdErr error, s *discordgo.Session, i *discordgo.InteractionCreate) {
	errorEmbed := errorEmbedMessage(cmdErr.Error())
	db.respondEmbed(errorEmbed, s, i)
}

func (db *DiscordBot) respondEmbed(embed *discordgo.MessageEmbed, s *discordgo.Session, i *discordgo.InteractionCreate) {
	response := &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{embed},
		},
	}

	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		log.Error("InteractionRespond error:", "error", err)
	}
}

func helpCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	embed := helpEmbed(s)
	db.respondEmbed(embed, s, i)
}

func claimCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	// Get the application command data
	data := i.ApplicationCommandData()

	// Extract the options
	testnetAddr := data.Options[0].StringValue()
	mainnetAddr := data.Options[1].StringValue()

	log.Info("new claim request", "discordID", i.Member.User.ID, "mainNetAddr", mainnetAddr, "testNetAddr", testnetAddr)

	command := fmt.Sprintf("claim %s %s %s", i.Member.User.ID, testnetAddr, mainnetAddr)

	result, err := db.BotEngine.Run(engine.AppIdDiscord, command, []string{testnetAddr, mainnetAddr})
	if err != nil {
		db.respondErrMsg(err, s, i)

		return
	}

	embed := claimEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func claimerInfoCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	testnetAddr := i.ApplicationCommandData().Options[0].StringValue()
	command := fmt.Sprintf("claimer-info %s", testnetAddr)

	result, err := db.BotEngine.Run(engine.AppIdDiscord, command, []string{testnetAddr})
	if err != nil {
		db.respondErrMsg(err, s, i)

		return
	}

	embed := claimerInfoEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func nodeInfoCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	valAddress := i.ApplicationCommandData().Options[0].StringValue()
	command := fmt.Sprintf("node-info %s", valAddress)

	result, err := db.BotEngine.Run(engine.AppIdDiscord, command, []string{valAddress})
	if err != nil {
		db.respondErrMsg(err, s, i)

		return
	}

	embed := nodeInfoEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func networkHealthCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	command := "network-health"

	result, err := db.BotEngine.Run(engine.AppIdDiscord, command, []string{})
	if err != nil {
		db.respondErrMsg(err, s, i)

		return
	}

	var color int
	if strings.Contains(result.Message, "Healthy") {
		color = GREEN
	} else {
		color = RED
	}

	embed := networkHealthEmbed(s, i, result, color)
	db.respondEmbed(embed, s, i)
}

func networkStatusCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	result, err := db.BotEngine.Run(engine.AppIdDiscord, "network", []string{})
	if err != nil {
		db.respondErrMsg(err, s, i)
		return
	}

	embed := networkStatusEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func walletCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	result, _ := db.BotEngine.Run(engine.AppIdDiscord, "wallet", []string{})

	embed := botWalletEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func claimStatusCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	result, _ := db.BotEngine.Run(engine.AppIdDiscord, "claim-status", []string{})

	embed := claimStatusEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func rewardCalcCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	stake := i.ApplicationCommandData().Options[0].StringValue()
	time := i.ApplicationCommandData().Options[1].StringValue()

	result, err := db.BotEngine.Run(engine.AppIdDiscord, fmt.Sprintf("calc-reward %v %v", stake, time), []string{stake, time})
	if err != nil {
		db.respondErrMsg(err, s, i)

		return
	}

	embed := rewardCalcEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func boosterPaymentCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	twitterName := i.ApplicationCommandData().Options[0].StringValue()
	valAddr := i.ApplicationCommandData().Options[1].StringValue()

	result, err := db.BotEngine.Run(engine.AppIdDiscord,
		fmt.Sprintf("booster-payment %v %v %v", i.Member.User.ID, twitterName, valAddr), []string{twitterName, valAddr})
	if err != nil {
		db.respondErrMsg(err, s, i)
		return
	}

	embed := boosterEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func boosterClaimCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	twitterID := i.ApplicationCommandData().Options[0].StringValue()

	result, err := db.BotEngine.Run(engine.AppIdDiscord, fmt.Sprintf("booster-claim %v", twitterID), []string{twitterID})
	if err != nil {
		db.respondErrMsg(err, s, i)
		return
	}

	embed := boosterEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func boosterWhitelistCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}

	twitterName := i.ApplicationCommandData().Options[0].StringValue()

	result, err := db.BotEngine.Run(engine.AppIdDiscord,
		fmt.Sprintf("booster-whitelist %v %v", twitterName, i.Member.User.ID), []string{twitterName})
	if err != nil {
		db.respondErrMsg(err, s, i)
		return
	}

	// Create a CommandResult instance with the message
	pubResult := engine.CommandResult{
		Message:    fmt.Sprintf("The Twitter account @%s has been successfully whitelisted!", twitterName),
		Successful: true, // Assuming the operation was successful
	}

	pubEmbed := boosterEmbed(s, i, &pubResult)
	_, err = s.ChannelMessageSendEmbed("1208143718482182184", pubEmbed)
	if err != nil {
		db.respondErrMsg(err, s, i)
		return
	}

	embed := boosterEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}

func boosterStatusCommandHandler(db *DiscordBot, s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !checkMessage(i, s, db.GuildID, i.Member.User.ID) {
		return
	}
	result, _ := db.BotEngine.Run(engine.AppIdDiscord, "booster-status", []string{})
	embed := boosterEmbed(s, i, result)
	db.respondEmbed(embed, s, i)
}
