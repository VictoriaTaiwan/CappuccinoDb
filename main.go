package main

import (
	"fmt"

	"cappuchinodb.com/main/app/data/models"
	"cappuchinodb.com/main/app/data/repositories"

	"github.com/gofiber/fiber/v2"
	"log"
	"os"
)

func main() {

	app := fiber.New()

	app.Get("/", indexHandler) // Add this
	app.Post("/", postHandler) // Add this 
	app.Put("/update", putHandler) // Add this 
	app.Delete("/delete", deleteHandler) // Add this

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))

	//testDatabase()

}

func indexHandler(c *fiber.Ctx) error {
	return c.SendString("Get")
}

 func postHandler(c *fiber.Ctx) error {
	return c.SendString("Post")
}

 func putHandler(c *fiber.Ctx) error {
	return c.SendString("Put")
}

 func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Delete")
}

func testDatabase(){
	sugarURL := "https://5.imimg.com/data5/AT/TN/MY-2/granulated-white-sugar-500x500.jpg"
	product := models.Product{
		Name:     "sugar",
		Calories: 48,
		UnitName: "table spoons",
		ImageSrc: sugarURL,
	}
	productCrErr := repositories.CreateProduct(product)
	if productCrErr == nil {
		fmt.Println("Sugar product was added to the database")
	} else {
		fmt.Println(productCrErr)
	}

	searchedProduct, searchedProductErr := repositories.GetProduct("sugar")
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

	recipeCrErr := repositories.CreateRecipe(pavlovaRecipe)
	if recipeCrErr == nil {
		fmt.Println("Pavlova cake was added to the database")
	} else {
		fmt.Println(recipeCrErr)
	}
}
