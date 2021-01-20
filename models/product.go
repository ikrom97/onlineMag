package models

import (
	"database/sql"
	"log"
	"onlineMag/db"
)

type Product struct {
	ID       int64  `json:"id"`
	Category string `json:"catalog_name"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Cost     int64  `json:"cost"`
	Status   string `json:"status"`
}
type RequestProductsList struct {
	CategoryName string `json:"category_name"`
}
type ResponseProduct struct {
	Name   string `json:"name"`
	Photo  string `json:"photo"`
	Cost   int64  `json:"cost"`
	Status string `json:"status"`
}
type AddProductResponse struct {
	Category string `json:"category"`
	Name     string `json:"name"`
	Photo    string `json:"photo"`
	Cost     int64  `json:"cost"`
	Status   string `json:"status"`
}

func ShowProductsByCategory(Db *sql.DB, category RequestProductsList) (productsList []ResponseProduct, err error) {
	row, err := Db.Query(db.ShowProductsByCategory, category.CategoryName)
	if err != nil {
		log.Println("Can't get products by category:",err)
		return
	}
	for row.Next() {
		var p ResponseProduct
		err := row.Scan(&p.Name, &p.Photo, &p.Cost, &p.Status)
		if err != nil {
			log.Println("Can't scan products for response:",err)
			continue
		}
		if p.Status == "В корзину" {
			productsList = append(productsList, p)
		}
	}
	return
}
func AddNewProduct(Db *sql.DB, product AddProductResponse) (err error) {
	_, err = Db.Exec(db.AddNewProduct, product.Category, product.Name, product.Photo, product.Cost, product.Status)
	if err != nil {
		log.Println("Can't exec new product:",err)
		return
	}
	return
}
func DeleteProduct(Db *sql.DB, name string) (err error) {
	_, err = Db.Exec(db.UpdateProductByName, "Removed", name)
	if err != nil {
		log.Println("Can't update products:",err)
		return
	}
	return
}
