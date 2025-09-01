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
	Steamid string `json:"steamid"`
}
