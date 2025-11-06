package models

// Recipe represents a single recipe class within the Classes array
type RawRecipe struct {
	ClassName                        string `json:"ClassName"`
	DisplayName                      string `json:"mDisplayName"`
	Ingredients                      string `json:"mIngredients"`
	Product                          string `json:"mProduct"`
	ManufacturingDuration            string `json:"mManufactoringDuration"`
	ProducedIn                       string `json:"mProducedIn"`
	VariablePowerConsumptionConstant string `json:"mVariablePowerConsumptionConstant"`
	VariablePowerConsumptionFactor   string `json:"mVariablePowerConsumptionFactor"`
}

type RecipeInfo struct {
	DisplayName                      string         `json:"displayName"`
	Ingredients                      map[string]int `json:"ingredients"`
	Product                          map[string]int `json:"product"`
	ManufacturingDuration            float64        `json:"manufactoringDuration"`
	ProducedIn                       string         `json:"producedIn"`
	VariablePowerConsumptionConstant float64        `json:"variablePowerConsumptionConstant"`
	VariablePowerConsumptionFactor   float64        `json:"variablePowerConsumptionFactor"`
}

type Recipe struct {
	ClassName map[string]RecipeInfo `json:"Recipe"`
}

// GameDataEntry represents each top-level object in the JSON array
type GameDataEntry struct {
	NativeClass string      `json:"NativeClass"`
	Classes     []RawRecipe `json:"Classes"`
}

type ItemRecipes struct {
	//map of product[recipeNames] with hash of docs used to generate
	Hash    string                `json:"Hash"`
	Item    map[string][]string   `json:"Product"`
	Recipes map[string]RecipeInfo `json:"Recipes"`
}
