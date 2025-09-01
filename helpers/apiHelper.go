package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type FriendListResponse struct {
	Friendlist struct {
		Friends []struct {
			Steamid string `json:"steamid"`
		} `json:"friends"`
	} `json:"friendlist"`
}

func getFriends(steamid string, APIKey string) ([]string, error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend", APIKey, steamid)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
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
