package jsonio

import (
	"maps"
	"slices"
	"sort"
)

type recipeAPI struct {
	data        ItemRecipes
	sortedItems []string
}

type RecipeAPI interface {
	Hash() string
	ItemRecipes(product string) []string
	RecipeDetails(recipeName string) RecipeInfo
	Products() []string
}

// go constructor?
func NewrecipeAPI(recipes ItemRecipes) RecipeAPI {
	items := slices.Collect(maps.Keys(recipes.Item))
	sort.Strings(items)
	return &recipeAPI{
		data:        recipes,
		sortedItems: items,
	}
}

// prob will get rid of this as its only used in main and not needed
func (api *recipeAPI) Hash() string {
	return api.data.Hash
}

func (api *recipeAPI) ItemRecipes(product string) []string {
	return api.data.Item[product]
}

func (api *recipeAPI) RecipeDetails(recipeName string) RecipeInfo {
	return api.data.Recipes[recipeName]
}

// sorted list of all Prodcuts(items)
func (api *recipeAPI) Products() []string {
	return api.sortedItems
}
