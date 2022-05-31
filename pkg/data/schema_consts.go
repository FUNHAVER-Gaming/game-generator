package data

const (
	//
	gameTable = "games"
	//Riot ID | Discord ID | Rating | Total Games Played | Last Game Played At | Won Rounds | Lost Rounds | Total Rounds | Game IDs Played |
	playerTable = "players"
	reportTable = "reports"

	createPlayerTable = "CREATE TABLE IF NOT EXISTS " + playerTable + " (" +
		"user_id VARCHAR(21), " +
		"riot_id VARCHAR(48), " +
		"riot_tag VARCHAR(24), " +
		"discord_id VARCHAR(18), " +
		"discord_username VARCHAR(32), " +
		"roles TEXT[], " +
		"rating SMALLINT, " +
		"total_games_played INTEGER, " +
		"last_game_played_timestamp TIMESTAMPTZ, " +
		"last_game_played_id VARCHAR(21), " +
		"won_rounds INTEGER, " +
		"lost_rounds INTEGER, " +
		"total_rounds INTEGER, " +
		"all_games_played TEXT[], " +
		"PRIMARY KEY (user_id, riot_id, discord_id)" +
		")"

	createPlayerGameDataType = "CREATE TYPE player_game_data AS (" +
		"user_id VARCHAR(21), " +
		"agent VARCHAR(16), " +
		"acs SMALLINT, " +
		"kills SMALLINT, " +
		"deaths SMALLINT, " +
		"assists SMALLINT, " +
		"first_bloods SMALLINT, " +
		"first_deaths SMALLINT, " +
		"plants SMALLINT, " +
		"defuses SMALLINT" +
		")"

	//Game ID | List of Players+Agents | Score | Map | Timestamp | Length | Team 1 Rating | Team 2 Rating | Team 1 Starting Side | Team 2 Starting Side | Winning Team |
	// Total Rounds Played |

	createGameTable = "CREATE TABLE IF NOT EXISTS " + gameTable + " (" +
		"game_id VARCHAR(21), " +
		"players player_game_data[], " +
		"score VARCHAR(5), " +
		"map VARCHAR(16), " +
		"game_start_time TIMESTAMPTZ," +
		"game_length BIGINT, " +
		"team_1_rating SMALLINT, " +
		"team_2_rating SMALLINT, " +
		//0 - def | 1 - offense
		"team_1_starting_side SMALLINT, " +
		"team_2_starting_side SMALLINT, " +
		"winning_team SMALLINT," +
		"total_rounds_played SMALLINT," +
		"PRIMARY KEY (game_id)" +
		")"

	getAllPlayers = "SELECT * FROM " + playerTable

	getPlayerFromDiscordId = "SELECT * FROM " + playerTable + " WHERE discord_id=$1"
	getPlayerFromRiotId    = "SELECT * FROM " + playerTable + " WHERE riot_id=$1"
	getPlayerFromUserId    = "SELECT * FROM " + playerTable + " WHERE user_id=$1"

	findUserIdFromRiotId    = "SELECT user_id FROM " + playerTable + " WHERE riot_id=$1"
	findUserIdFromDiscordId = "SELECT user_id FROM " + playerTable + " WHERE discord_id=$1"

	// savePlayer Insert a player into the player table: UserID, RiotID, DiscordID, DiscordUsername, Roles, Rating
	savePlayer = "INSERT INTO " + playerTable + " (user_id, riot_id, riot_tag, discord_id, discord_username, roles, rating)" +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"

	updateRoles  = "UPDATE + " + playerTable + " SET roles=$1 WHERE user_id=$2"
	updateRating = "UPDATE + " + playerTable + " SET rating=$1 WHERE user_id=$2"

	// playerPlayGame Update a players rating, total games played, last game timestamp and ID, won rounds and lost rounds, for a given user ID
	playerPlayGame = "UPDATE " + playerTable + " SET rating=$1, " +
		"total_games_played=total_games_played+1, " +
		"last_game_played_timestamp=$2, " +
		"last_game_played_id=$3, " +
		"won_rounds=won_rounds+$4, " +
		"lost_rounds=lost_rounds+$4, " +
		"total_rounds=(lost_rounds+won_rounds)" +
		"WHERE user_id=$1"

	// createGame Insert a new game into the DB: GameID, Players, Map, GameStartTime, T1Rating, T2Rating, T1StartingSide, T2StartingSide
	createGame = "INSERT INTO " + gameTable + " (game_id, map, game_start_time, team_1_rating, team_2_rating, team_1_starting_side, team_2_starting_side) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)"

	// getGameStartTime Get the timestamp of the games starting time: GameID
	getGameStartTime = "SELECT game_start_time FROM " + gameTable + " WHERE $game_id=$1"

	// finishGame Update the game for the given id: GameLength, Score, WinningTeam, TotalRoundsPlayed, GameID
	finishGame = "UPDATE " + gameTable + " SET game_length=$1, score=$2, winning_team=$3, total_rounds_played=$4 WHERE game_id=$5"
)
