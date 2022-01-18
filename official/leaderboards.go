package official

type Region int

// Region constants used by the api.
const (
	RegionEurope       Region = 0
	RegionMiddleEast   Region = 1
	RegionAsia         Region = 2
	RegionNorthAmerica Region = 3
	RegionSouthAmerica Region = 4
	RegionOceana       Region = 5
	RegionAfrica       Region = 6
	RegionGlobal       Region = 7
)

type TeamSize string

// TeamSize constants used by the api.
const (
	TeamSizeAll TeamSize = ""
	TeamSize1v1 TeamSize = "1v1"
	TeamSize2v2 TeamSize = "2v2"
	TeamSize3v3 TeamSize = "3v3"
	TeamSize4v4 TeamSize = "4v4"
)

type MatchType string

// MatchType constants used by the api.
const (
	MatchTypeUnranked       MatchType = "unranked"
	MatchTypeCustom         MatchType = "custom"
	MatchTypeAIEasy         MatchType = "aieasy"
	MatchTypeAIIntermediate MatchType = "aimedium"
	MatchTypeAIHard         MatchType = "aihard"
	MatchTypeAIExpert       MatchType = "aiexpert"
)

type Versus string

const (
	VersusAll     Versus = ""
	VersusPlayers Versus = "players"
	VersusAI      Versus = "ai"
)

// LeaderboardArgs are used to query the leaderboard api.
type LeaderboardArgs struct {
	Region       Region    `json:"region"`
	Versus       Versus    `json:"versus"`
	MatchType    MatchType `json:"matchType"`
	TeamSize     TeamSize  `json:"teamSize"`
	SearchPlayer string    `json:"searchPlayer"`
	Page         int       `json:"page"`
	Count        int       `json:"count"`
}

// Leaderboard is returned from the official leaderboard
// api call. IT has a count of all players as well as a paged
// list of
type Leaderboard struct {
	Count int `json:"count"`
	Items []struct {
		GameID       interface{} `json:"gameId"`
		UserID       interface{} `json:"userId"`
		RlUserID     int         `json:"rlUserId"`
		UserName     string      `json:"userName"`
		AvatarURL    string      `json:"avatarUrl"`
		PlayerNumber interface{} `json:"playerNumber"`
		Elo          int         `json:"elo"`
		EloRating    int         `json:"eloRating"`
		Rank         int         `json:"rank"`
		Region       string      `json:"region"`
		Wins         int         `json:"wins"`
		WinPercent   float64     `json:"winPercent"`
		Losses       int         `json:"losses"`
		WinStreak    int         `json:"winStreak"`
	} `json:"items"`
}
