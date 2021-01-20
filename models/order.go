package models

import (
	"database/sql"
	"log"
	"onlineMag/db"
	"time"
)

type Order struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	OrderedDate string `json:"ordered_date"`
	UpdatedDate string `json:"updated_date"`
	ProductName string `json:"product_name"`
	Address     string `json:"address"`
	Status      string `json:"status"`
}
type OrderRequest struct {
	ProductName string `json:"product_name"`
	Address     string `json:"address"`
	Status      string `json:"status"`
}
type OrderBody struct {
	ProductName string `json:"product_name"`
	Address string `json:"address"`
}
type OrdersID struct {
	OrdersID int64 `json:"orders_id"`
}

func AddNewOrder(Db *sql.DB, id int64, body OrderBody) (err error) {
	Time := time.Now().Format("02-Jan-2006 15:41")
	status := "Заказано"
	_, err = Db.Exec(db.AddNewOrder, id, Time, body.ProductName, body.Address, status)
	if err != nil {
		log.Println("Can't add new order:",err)
		return
	}
	return
}

func ShowOrders(Db *sql.DB, id int64) (orderList []OrderRequest, err error) {
	var list []OrderRequest
	row, err := Db.Query(db.GetOrdersByUserID, id)
	if err != nil {
		log.Println("Can't get product by userID:",err)
		return
	}
	for row.Next() {
		var p OrderRequest
		err := row.Scan(&p.ProductName, &p.Address, &p.Status)
		if err != nil {
			log.Println("Can't scan products:",err)
			continue
		}
		list = append(list, p)
	}
	for i := 0; i < len(list); i++ {
		if list[i].Status == "Заказано" {
			orderList = append(orderList, list[i])
		}
	}
	return
}
func CompleteOrder(Db *sql.DB, ID int64) (err error) {
	Time := time.Now().Format("02-Jan-2006 15:41")
	status := "Выполнен"
	_, err = Db.Exec(db.UpdateOrderById, Time, status, ID)
	if err != nil {
		log.Println("Can't update order:", err)
		return
	}
	return
}
func GetOrdersByStatus(Db *sql.DB, status string) ( list []Order,err error) {
	rows, err := Db.Query(db.GetOrdersByStatus, status)
	if err != nil {
		log.Println("Can't select orders by status:", err)
		return
	}
	for rows.Next() {
		var p Order
		err := rows.Scan(&p.ID, &p.UserID, &p.OrderedDate, &p.UpdatedDate, &p.ProductName, &p.Address, &p.Status)
		if err != nil {
			log.Println("Can't scan orders:", err)
			continue
		}
		list = append(list, p)
	}
	return
}
func RemoveOrder(Db *sql.DB, id int64, body OrderBody) (err error) {
	Time := time.Now().Format("02-Jan-2006 15:41")
	status := "Отменен"
	_, err = Db.Exec(db.UpdateOrdersStatus, Time, status, id, body.ProductName)
	if err != nil {
		log.Println("Can't update order:",err)
		return
	}
	return
}