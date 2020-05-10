package model

import "fmt"

// Ingredient is what we use in the recipe
type Ingredient struct {
	Amount string `as:"amount"`
	Type   string `as:"type"`
}

// Recipe contains the title, ingredients and process of an Recipe
type Recipe struct {
	Title       string       `as:"title"`
	Ingredients []Ingredient `as:"ingredients"`
	Process     string       `as:"process"`
}

// AddRecipe adds a recipe and gets back an error is something went wrong
func (db *DB) AddRecipe(title string, ingredients []Ingredient, process string) error {
	// Add title and content into aerospike
	k, err := NewKey(title)
	if err != nil {
		return err
	}

	r := &Recipe{
		Title:       title,
		Ingredients: ingredients,
		Process:     process,
	}

	db.PutObject(nil, k, r)
	return nil
}

// GetAllRecipes retrieves all recipe titles
func (db *DB) GetAllRecipes() ([]string, error) {
	rcs, err := db.ScanAll(nil, Database, "recipes")
	if err != nil {
		return nil, err
	}

	var list []string
	for r := range rcs.Results() {
		recipe := r.Record.Bins["title"]

		list = append(list, fmt.Sprintf("%v", recipe))
	}

	return list, nil
}

// GetRecipe takes an recipe id and returns the recipe or throws an error if recipe isn't found
func (db *DB) GetRecipe(title string) (*Recipe, error) {
	key, err := NewKey(title)
	if err != nil {
		return nil, err
	}

	r := &Recipe{}
	err = db.GetObject(nil, key, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
