package main

import (
	"Gamebanfinder/helpers"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	steamid := ""
	for {
		APIKEY := os.Getenv("APIKEY")
		println("Welcome to Gamebanfinder import the steamid below")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			steamid = input.Text()
			if len(steamid) != 17 {
				fmt.Println("Invalid steamid should be 17 numbers long")
				break
			}
			Friends, err := helpers.GetFriends(steamid, APIKEY)
			if err != nil {
				fmt.Println(err)
				break
			}
			if len(Friends) == 0 {
				fmt.Printf("No Friends found acount could be private steamid: %v\n", steamid)
			}
			bans, err := helpers.CheckBans(Friends, APIKEY)
			if err != nil {
				fmt.Println(err)
				break
			}
			if len(bans) == 0 {
				fmt.Println("No bans found in friends from steamid:" + steamid)
			}
			for _, ban := range bans {
				if ban.VACBanned {
					if ban.NumberOfGameBans <= 0 || ban.SteamId != "" {
						println("either this acount hase no gamebans or the steamid returned from steam is empty")
					}
					isBannedString := "yes"
					fmt.Println("steamid" + ban.SteamId)
					fmt.Println("Vacbanned" + isBannedString)
					numOfGameBans := strconv.Itoa(ban.NumberOfGameBans)
					fmt.Println("Number of vacbans" + numOfGameBans)
					numOfDaysSinceLastBan := strconv.Itoa(ban.DaysSinceLastBan)
					fmt.Println("DaysSinceLastBan" + numOfDaysSinceLastBan)
				}
			}
		}
	}
}
