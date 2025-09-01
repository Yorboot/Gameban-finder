package Gamebanfinder

import (
	//"Gamebanfinder/helpers"
	"bufio"
	"os"
)

func main() {
	steamid := ""
	for {
		println("Welcome to Gamebanfinder import the steamid below")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			steamid = input.Text()
			println(steamid)
		}
	}
	//Friends, err = getFriends()
}
