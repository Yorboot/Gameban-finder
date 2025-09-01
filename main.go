package main

import (
	//"Gamebanfinder/helpers"
	"bufio"
	"fmt"
	"os"
)

func main() {
	steamid := ""
	for {
		println("Welcome to Gamebanfinder import the steamid below")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			steamid = input.Text()
			if len(steamid) != 17 {
				fmt.Errorf("invalid steamid should be 17 characters long")
			}
			println(len(steamid))
		}
	}
	//Friends, err = getFriends()
}
