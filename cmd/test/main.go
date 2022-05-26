package main

func main() {
	/*dg, err := discordgo.New("Bot " + service.DiscordToken)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	team1 := []string{"vox", "jovi", "mini", "ion", "tim"}

	team1msg := "Team 1: "
	for i, user := range team1 {
		team1msg += user
		if i != len(team1)-1 {
			team1msg += ", "
		}
	}

	team2msg := "Team 2: "
	for i, user := range team1 {
		team2msg += "NOT-" + user
		if i != len(team1)-1 {
			team2msg += ", "
		}
	}

	attackingMsg := ""
	r := rand.New(rand.NewSource(time.Now().Unix()))

	if r.Intn(100) >= 50 {
		attackingMsg = "Team 1 is attack"
	} else {
		attackingMsg = "Team 2 is attack"
	}

	lobbyLeader := team1[r.Intn(len(team1)-1)]

	fields := []*discordgo.MessageEmbedField{
		{Name: "Team 1", Value: team1msg, Inline: true},
		{Name: "Team 2", Value: team2msg},
		{Name: "Attacking", Value: attackingMsg},
		{Name: "Lobby Leader", Value: lobbyLeader},
		{Name: "ID", Value: fmt.Sprintf("||%v||", uuid.NewString())},
	}

	embed := &discordgo.MessageEmbed{
		URL:         "http://localhost:8000",
		Type:        discordgo.EmbedTypeRich,
		Title:       "game-1: Game Created",
		Description: "New FUNHAVER Gaming Ranked Game",
		Timestamp:   time.Now().Format(time.RFC3339),
		Color:       24,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "The FUNHAVER Bot",
		},
		Fields: fields,
	}

	_, err = dg.ChannelMessageSendEmbed(service.Game1ID, embed)
	resp, err := dg.ChannelMessages(service.Game1ID, 100, "", "", "")

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var ids []string

	for _, msg := range resp {
		ids = append(ids, msg.ID)
	}

	err = dg.ChannelMessagesBulkDelete(service.Game1ID, ids)*/

}
