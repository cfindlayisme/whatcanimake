package model

type Recipes struct {
	Recipes []Recipe `json:"recipes"`
}

type Recipe struct {
	Name        string `json:"name"`
	Ingredients []struct {
		Ingredient string `json:"ingredient"`
		Quantity   string `json:"quantity"`
	} `json:"ingredients"`
	Instructions []string `json:"instructions"`
}

type QueryParameters struct {
	Count         *int    `json:"Count"`
	Ingredients   *string `json:"Ingredients"`
	PantryStaples *bool   `json:"PantryStaples"`
}

const PantryStaples = "flour, sugar, vegetable oil, baking powder, sugar, salt"
