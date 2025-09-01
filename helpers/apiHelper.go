package helpers

type FriendListResponse struct {
	Friendlist struct {
		Friends []struct {
			Steamid string `json:"steamid"`
		} `json:"friends"`
	} `json:"friendlist"`
}
