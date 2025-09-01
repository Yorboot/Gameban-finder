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
				fmt.Println("Invalid steamid should be 17 numbers long")
				break
			}
			println(len(steamid))
		}
	}
	//Friends, err = getFriends()
}
