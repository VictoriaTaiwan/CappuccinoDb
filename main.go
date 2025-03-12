package main

import (
	"fmt"

	"cappuchinodb.com/main/app/router"
	"cappuchinodb.com/main/app/pkg/handlers"
	"cappuchinodb.com/main/app/pkg/models"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	router.SetupRoutes(app)

	app.Listen(":8080")

	//testDatabase()
}

func testDatabase() {
	sugarURL := "https://5.imimg.com/data5/AT/TN/MY-2/granulated-white-sugar-500x500.jpg"
	product := models.Product{
		Name:     "sugar",
		Calories: 48,
		UnitName: "table spoons",
		ImageSrc: sugarURL,
	}
	productCrErr := handlers.CreateProduct(product)
	if productCrErr == nil {
		fmt.Println("Sugar product was added to the database")
	} else {
		fmt.Println(productCrErr)
	}

	searchedProduct, searchedProductErr := handlers.GetProduct("sugar")
	if searchedProductErr == nil {
		fmt.Println("Sugar is counted by " + searchedProduct.UnitName)
	} else {
		fmt.Println(searchedProductErr)
	}

	pavlovaURL := "https://www.sugarsaltmagic.com/wp-content/uploads/2024/11/The-Perfect-Pavlova-Recipe-2FEAT-500x375.jpg"
	recipeInfo := models.Product{
		Name:     "Pavlova cake",
		Calories: 246,
		UnitName: "1 serving",
		ImageSrc: pavlovaURL,
	}
	ingredients := make(map[int]string)
	ingredients[0] = "200 gram"
	instructions := "In a stand mixer fit with the whisk attachment, beat the egg whites" +
		"on low speed until their foamy and little bubbles form."
	pavlovaRecipe := models.Recipe{
		Information:  recipeInfo,
		Ingredients:  ingredients,
		Instructions: instructions,
	}

	recipeCrErr := handlers.CreateRecipe(pavlovaRecipe)
	if recipeCrErr == nil {
		fmt.Println("Pavlova cake was added to the database")
	} else {
		fmt.Println(recipeCrErr)
	}
}
