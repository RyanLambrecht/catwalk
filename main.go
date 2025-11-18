package main

import (
	"fmt"
	"os"

	"github.com/RyanLambrecht/catwalk/jsonio"
	"github.com/RyanLambrecht/catwalk/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	api, err := jsonio.InitalizeAPI()
	if err != nil {
		panic(fmt.Sprintf("API startup failed: %v\n", err))

	}
	tempItems := api.Products()

	//for _, i := range tempItem {
	//	fmt.Println(i)
	//}

	fmt.Println(api.ItemRecipes(tempItems[1]))

	p := tea.NewProgram(ui.NewAppModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("error: %v", err)
		os.Exit(1)
	}
}

//if err != nil {
//		fmt.Fprintf(os.Stderr, "Startup failed: %v\n", err)
//		os.Exit(1)
//	}
//
//	fmt.Println(api.data.Hash)
//	product := "Desc_AlienProtein_C"
//	productRecipes := api.GetProductRecipes(product)
//	fmt.Println(len(productRecipes))
//
//	for _, recipe := range productRecipes {
//		fmt.Println(recipe)
//	}
