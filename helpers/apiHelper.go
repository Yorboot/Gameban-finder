package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
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

const (
	APIKEY = ""
)

func getFriends(steamid string, APIKey string) ([]string, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend", APIKey, steamid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	var data FriendListResponse
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}
	friends := []string{}
	for _, friend := range data.Friendlist.Friends {
		friends = append(friends, friend.Steamid)
	}
	return friends, nil
}
func checkBans(friendids []string, APIKey string) ([]PlayerBan, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetPlayerBans/v1/?key=%s&steamids=%s", APIKey, strings.Join(friendids, ","))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, _ := io.ReadAll(resp.Body)
	var bans PlayerBansResponse
	if err := json.Unmarshal(body, &bans); err != nil {
		return nil, err
	}
	return bans.Players, nil
}
