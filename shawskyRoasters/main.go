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

	var storeMenu = menu.ImportMenu()

loop:
	for {
		fmt.Println("Welcome to Shawsky Roasters! Please select an option to continue...")
		fmt.Println("1.) Print Menu")
		fmt.Println("2.) Add menu item")
		fmt.Println("Q.) Quit")
		choice, _ := userInput.ReadString('\n')

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintMenu(&storeMenu)
		case "2":
			name, category, prices := ReadMenuItem()
			err := menu.AddMenuItem(&storeMenu.MenuItems, name, category, prices)
			if err != nil {
				fmt.Println(err)
				fmt.Println("Unable to create new menu item. Please try again.")
				fmt.Println()
			}
		case "q":
		case "Q":
			break loop
		default:
			fmt.Println("Unknown option... Please try again")
		}
	}
}

func ReadMenuItem() (string, string, []string) {
	fmt.Println("Please enter the category of the new item")
	category, _ := userInput.ReadString('\n')
	category = strings.TrimSpace(category)

	fmt.Println("Please enter the name of the new item")
	name, _ := userInput.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Please enter the price of a small, medium, and large item delimited by pipes")
	prices, _ := userInput.ReadString('\n')
	prices = strings.TrimSpace(prices)
	priceList := strings.Split(prices, "|")

	return name, category, priceList
}
