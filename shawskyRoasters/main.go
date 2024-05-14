package main

import (
	"bufio"
	"fmt"
	"os"
	"shawskyRoasters/menu"
	"strings"
)

var userInput = bufio.NewReader(os.Stdin)

func main() {

loop:
	for {
		fmt.Println("Welcome to Shawsky Roasters! Please select an option to continue...")
		fmt.Println("1.) Print Menu")
		fmt.Println("2.) Add menu item")
		fmt.Println("Q.) Quit")
		choice, _ := userInput.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.ImportMenu()
		case "q":
			break loop
		case "Q":
			break loop
		default:
			fmt.Println("Unknown option... Please try again")
		}

	}

}
