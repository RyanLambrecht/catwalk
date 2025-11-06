package jsonio

import (
	"maps"
	"slices"

	"github.com/RyanLambrecht/catwalk/models"
)

type RecipeAPI struct {
	data models.ItemRecipes
}

// go constructor?
func NewRecipeAPI(recipes models.ItemRecipes) *RecipeAPI {
	return &RecipeAPI{
		data: recipes,
	}
}

// prob will get rid of this
func (api *RecipeAPI) GetHash() string {
	return api.data.Hash
}

func (api *RecipeAPI) GetItemRecipes(product string) []string {
	return api.data.Item[product]
}

func (api *RecipeAPI) GetRecipeDetails(recipeName string) models.RecipeInfo {
	return api.data.Recipes[recipeName]
}

func (api *RecipeAPI) GetItems() []string {
	return slices.Collect(maps.Keys(api.data.Item))
}
