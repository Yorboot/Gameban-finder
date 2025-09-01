package Gamebanfinder

const (
	APIKEY = ""
)

type PlayerBan struct {
	SteamId          string `json:"steamId"`
	CommunityBanned  bool   `json:"community_banned"`
	VACBanned        bool   `json:"vac_banned"`
	NumberOfVACBans  int    `json:"NumberOfVACBans"`
	NumberOfGameBans int    `json:"NumberOfGameBans"`
	DaysSinceLastBan int    `json:"DaysSinceLastBan"`
	EconomyBan       string `json:"EconomyBan"`
}
type PlayerBansResponse struct {
	Players []PlayerBan `json:"players"`
}
