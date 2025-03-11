package models

type Product struct {
	ID       int     `json:"ID"`
	Name     string  `json:"name"`
	Calories float64 `json:"calories"` // per 1 unit
	UnitName string  `json:"unit_name"` // tsp, g, ml etc
	ImageSrc string  `json:"image_src"`
}

type Recipe struct {
	Information  Product         `json:"information"` // Main information
	Ingredients  map[int]string `json:"ingredients"`// productID - description
	Instructions string          `json:"instructions"` // Put 2 spoons of sugar to the bowl
}
