package menu

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Price struct {
	Small  float64 `json:"small"`
	Medium float64 `json:"medium"`
	Large  float64 `json:"large"`
}

type MenuItem struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Prices   Price  `json:"prices"`
}

type Menu struct {
	MenuItems []MenuItem
}

//go:embed menu.json
var fileContents embed.FS

func ImportMenu() Menu {
	config, _ := fileContents.ReadFile("menu.json")

	var menu Menu
	menu.MenuItems = []MenuItem{}

	if err := json.Unmarshal(config, &menu.MenuItems); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return menu
	}
	return menu
}

func PrintMenu(input *Menu) {
	fmt.Println()
	for _, item := range input.MenuItems {
		fmt.Printf("Name: %s, Category: %s \n", item.Name, item.Category)
		fmt.Printf("\tPricing => Small: $%.2f    Medium: $%.2f    Large: $%.2f  \n", item.Prices.Small, item.Prices.Medium, item.Prices.Large)
		fmt.Println()
	}
}

func AddMenuItem(items *[]MenuItem, name string, category string, prices []string) error {
	for _, item := range *items {
		if item.Name == name {
			return errors.New("The menu item already exists; please try again")
		}
	}

	if len(prices) != 3 {
		return errors.New("Unable to detect three distinct prices for the new menu item")
	}

	smallPrice, e := strconv.ParseFloat(strings.TrimSpace(prices[0]), 64)
	if e != nil {
		return errors.New("Unable to parse the small price provided")
	}
	mediumPrice, e := strconv.ParseFloat(strings.TrimSpace(prices[1]), 64)
	if e != nil {
		return errors.New("Unable to parse the medium price provided")
	}

	largePrice, e := strconv.ParseFloat(strings.TrimSpace(prices[2]), 64)
	if e != nil {
		return errors.New("Unable to parse the medium price provided")
	}

	newItem := MenuItem{Name: name, Category: category, Prices: Price{Small: smallPrice, Medium: mediumPrice, Large: largePrice}}

	*items = append(*items, newItem)
	return nil
}
