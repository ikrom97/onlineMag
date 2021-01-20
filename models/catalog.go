package models

import (
	"database/sql"
	"log"
	"onlineMag/db"
)

type Catalog struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Remove bool   `json:"remove"`
}

func GetCatalogList(Db *sql.DB) (catalogList []Catalog, err error) {
	row, err := Db.Query(db.GetCatalogList)
	if err != nil {
		log.Println("Can't get catalog:", err)
		return
	}
	for row.Next() {
		var p Catalog
		err := row.Scan(&p.ID, &p.Name, &p.Remove)
		if err != nil {
			log.Println("Can't scan catalog:",err)
			continue
		}
		if !p.Remove {
			catalogList = append(catalogList, p)
		}
	}
	return
}
func CheckHasCategory(Db *sql.DB, name string) (category Catalog, err error) {
	row := Db.QueryRow(db.GetCatalogByName, name)
	err = row.Scan(&category.ID, &category.Name, &category.Remove)
	if err != nil {
		log.Println("Can't scan category:",err)
		return
	}
	return
}
func AddCategory(Db *sql.DB, catalog Catalog) (err error) {
	_, err = Db.Exec(db.UpdateCategoryByID, catalog.ID)
	if err != nil {
		log.Println("Can't update category:",err)
		return
	}
	return
}
func AddNewCategory(Db *sql.DB, name string) (err error) {
	_, err = Db.Exec(db.AddNewCategory, name)
	if err != nil {
		log.Println("Can't add new category:",err)
		return
	}
	return
}
func DeleteCategory(Db *sql.DB, name string) (err error) {
	_, err = Db.Exec(db.UpdateCategoryStatus, true, name)
	if err != nil {
		log.Println("Can't update category:",err)
		return
	}
	return
}
