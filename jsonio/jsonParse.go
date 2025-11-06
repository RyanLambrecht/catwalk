package jsonio

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/RyanLambrecht/catwalk/models"
)

const RecipesOutputFile = "lib/recipes.json"

var itemRegexp = regexp.MustCompile(`(Desc_\w+_C)'[^,]*,Amount=(\d+)`)
var ProducedInRegExp = regexp.MustCompile(`\.(\w+_C)`)

func refineRecipe(rawRecipe models.RawRecipe) models.RecipeInfo {
	var recipe models.RecipeInfo
	var err error

	recipe.DisplayName = rawRecipe.DisplayName

	recipe.ManufacturingDuration, err = strconv.ParseFloat(rawRecipe.ManufacturingDuration, 64)
	if err != nil {
		fmt.Printf("error parsing ManufacturingDuration: %v\n", err)
		os.Exit(1)
	}
	recipe.VariablePowerConsumptionConstant, err = strconv.ParseFloat(rawRecipe.VariablePowerConsumptionConstant, 64)
	if err != nil {
		fmt.Printf("error parsing VariablePowerConsumptionConstant: %v\n", err)
		os.Exit(1)
	}
	recipe.VariablePowerConsumptionFactor, err = strconv.ParseFloat(rawRecipe.VariablePowerConsumptionFactor, 64)
	if err != nil {
		fmt.Printf("error parsing VariablePowerConsumptionFactor: %v\n", err)
		os.Exit(1)
	}

	ingredients := make(map[string]int)
	for _, match := range itemRegexp.FindAllStringSubmatch(rawRecipe.Ingredients, -1) {
		ingredients[match[1]], err = strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("error parsing ingredient amount: %v\n", err)
			os.Exit(1)
		}
	}
	recipe.Ingredients = ingredients

	product := make(map[string]int)
	for _, match := range itemRegexp.FindAllStringSubmatch(rawRecipe.Product, -1) {
		product[match[1]], err = strconv.Atoi(match[2])
		if err != nil {
			fmt.Printf("error parsing product amount: %v\n", err)
			os.Exit(1)
		}
	}
	recipe.Product = product

	recipe.ProducedIn = ProducedInRegExp.FindStringSubmatch(rawRecipe.ProducedIn)[1]

	return recipe
}

func filterRecipes(entries []models.GameDataEntry, hash string) models.ItemRecipes {
	filteredRecipes := models.ItemRecipes{Hash: hash, Item: make(map[string][]string), Recipes: make(map[string]models.RecipeInfo)}

	for _, entry := range entries {
		for _, rawRecipe := range entry.Classes {
			if strings.Contains(rawRecipe.ProducedIn, "/Factory/") {
				recipe := refineRecipe(rawRecipe)
				filteredRecipes.Recipes[rawRecipe.ClassName] = recipe
				for product := range recipe.Product {
					filteredRecipes.Item[product] = append(filteredRecipes.Item[product], rawRecipe.ClassName)
				}

			}
		}
	}

	return filteredRecipes
}

func exportJson[T any](data T, fileName string) {
	jsonData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		fmt.Printf("Error marshalling data: %v\n", err)
		return
	}

	// Ensure lib directory exists
	dir := filepath.Dir(fileName)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Printf("Error creating directory: %v\n", err)
		return
	}

	err = os.WriteFile(fileName, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error creating json: %v\n", err)
	}
}

// returns true if the recipe JSON is up to date
func checkExistingRecipes(filename, currentHash string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false // File doesn't exist, need to process
	}

	// File exists, load and check hash
	existingRecipes, err := LoadJsonFile[models.ItemRecipes](filename)
	if err != nil {
		fmt.Printf("Error loading existing recipes: %v\n", err)
		return false
	}

	// Compare hashes
	if existingRecipes.Hash == currentHash {
		fmt.Println("Existing JSON is up to date")
		return true
	}

	fmt.Println("Source data has changed, updating recipes...")
	return false
}

// LoadJsonFile loads JSON data from a file into the provided destination
func LoadJsonFile[T any](filepath string) (T, error) {
	var result T

	byteValue, err := os.ReadFile(filepath)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func JsonParse() (string, error) {
	var filePath string
	if len(os.Args) == 2 {
		filePath = os.Args[1]
		fmt.Printf("Parsing %s\n", filePath)
	} else {
		filePath = filepath.Join("lib", "en-US.json")
		fmt.Printf("No path provided, parsing default path: %s\n", filePath)
	}

	// Read and hash the file for change detection
	fileData, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return "", err
	}

	byteHash := sha256.Sum256(fileData)
	hash := hex.EncodeToString(byteHash[:])

	if checkExistingRecipes(RecipesOutputFile, hash) {
		return RecipesOutputFile, nil
	}

	var allEntries []models.GameDataEntry
	allEntries, err = LoadJsonFile[[]models.GameDataEntry](filePath)
	if err != nil {
		fmt.Printf("Error loading JSON: %v\n", err)
		return "", err
	}

	// Find recipe entries
	var recipeEntries []models.GameDataEntry
	for _, entry := range allEntries {
		if strings.Contains(entry.NativeClass, "FGRecipe") {
			recipeEntries = append(recipeEntries, entry)
			break
		}
	}

	filteredRecipes := filterRecipes(recipeEntries, hash)
	exportJson(filteredRecipes, RecipesOutputFile)

	return RecipesOutputFile, nil
}
