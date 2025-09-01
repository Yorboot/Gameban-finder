package Gamebanfinder

const (
	APIKEY = ""
)

type FriendListResponse struct {
	Friendlist struct {
		Friends []struct {
			Steamid string `json:"steamid"`
		} `json:"friends"`
	} `json:"friendlist"`
}
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
