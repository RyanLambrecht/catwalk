package main

import (
	"fmt"
	"os"

	"github.com/RyanLambrecht/catwalk/jsonio"
	"github.com/RyanLambrecht/catwalk/models"
)

func startupAPI() (*jsonio.RecipeAPI, error) {
	recipeFileLocation, err := jsonio.JsonParse()
	if err != nil {
		return nil, fmt.Errorf("error parsing JSON: %w", err)
	}

	itemRecipes, err := jsonio.LoadJsonFile[models.ItemRecipes](recipeFileLocation)
	if err != nil {
		return nil, fmt.Errorf("error loading JSON: %w", err)
	}

	api := jsonio.NewRecipeAPI(itemRecipes)
	return api, nil
}

func main() {
	api, err := startupAPI()
	if err != nil {
		fmt.Fprintf(os.Stderr, "startup failed: %v\n", err)
		os.Exit(1)

	}
	fmt.Printf("api hash: %v\n", api.GetHash())

	tempItems := api.GetItems()

	//for _, i := range tempItem {
	//	fmt.Println(i)
	//}

	fmt.Println(api.GetItemRecipes(tempItems[1]))
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
