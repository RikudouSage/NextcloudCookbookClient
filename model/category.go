package model

type Category struct {
	Name        string `json:"name"`
	RecipeCount uint   `json:"recipe_count"`
}
