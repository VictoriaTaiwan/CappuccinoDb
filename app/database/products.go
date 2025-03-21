package database

import (
	"context"
	"cappuchinodb.com/main/app/pkg/models"
)

func CreateProduct(product models.Product) error {
	_, err := DB.Exec(context.Background(),
		"INSERT INTO products (name, calories, unit_name, image_src) VALUES ($1, $2, $3, $4)",
		product.Name, product.Calories, product.UnitName, product.ImageSrc)
	if err != nil {
		return err
	}

	return nil
}

// Get product by name
func GetProduct(name string) (models.Product, error) {
	var product models.Product
	row := DB.QueryRow(context.Background(),
		"SELECT id, name, calories, unit_name, image_src FROM products WHERE name=$1",
		name)
	if err := row.Scan(&product.ID, &product.Name, &product.Calories, &product.UnitName, &product.ImageSrc); err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// Update recipe's ingredients by recipe's ID
func UpdateProduct(productID int, product models.Product) error {
	_, err := DB.Exec(context.Background(),
		"UPDATE products SET name=$1, calories=$2, unit_name=$3, image_src=$4 WHERE id=$5",
		product.Name, product.Calories, product.UnitName, product.ImageSrc, productID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(productID int) error {
	_, err := DB.Exec(context.Background(), "DELETE FROM products WHERE id=$1", productID)
	if err != nil {
		return err
	}

	return nil
}
