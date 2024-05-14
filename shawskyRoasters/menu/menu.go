package menu

import (
	"embed"
	"encoding/json"
	"fmt"
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

func ImportMenu() {
	config, _ := fileContents.ReadFile("menu.json")

	var menu []MenuItem

	if err := json.Unmarshal(config, &menu); err != nil {
		fmt.Printf("failed to unmarshal json file, error: %v", err)
		return
	}
	for _, item := range menu {
		fmt.Printf("Name: %s, Category: %s \n", item.Name, item.Category)
		fmt.Printf("\tPricing => Small: $%.2f    Medium: $%.2f    Large: $%.2f  \n", item.Prices.Small, item.Prices.Medium, item.Prices.Large)
	}
}
