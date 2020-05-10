package model

import (
	as "github.com/aerospike/aerospike-client-go"
)

// RecipeStore implements the retrieval and writes of recipes to the chosen database
type RecipeStore interface {
	AddRecipe(string, []Ingredient, string) error
	GetAllRecipes() ([]string, error)
	GetRecipe(string) (*Recipe, error)
}

//DB Is the database handle
type DB struct {
	*as.Client
}

// Database is the database recipes are stored in
var Database string = "recipes"

// NewDB opens up a new connection to the chosen database
func NewDB(host string, port int) (*DB, error) {
	asClient, err := as.NewClient(host, port)
	if err != nil {
		return nil, err
	}

	return &DB{asClient}, nil
}

// NewKey wraps around aerospike's NewKey
func NewKey(key string) (*as.Key, error) {
	return as.NewKey(Database, "recipes", key)
}

// NewBin wraps around aerospike's NewBin
func NewBin(key string, value interface{}) *as.Bin {
	return as.NewBin(key, value)
}
