package app

import (
	"fmt"
	"log"
	"net/http"
	"onlineMag/middlewares"
)

func (server *MainServer) InitRoutes() {
	fmt.Println("Routes are init in localhost:8888")

	server.router.GET("/", server.MainHandler)
	server.router.GET("/api/Sign-in", server.SignInHandler)
	server.router.GET("/api/catalog", server.CatalogHandler)
	server.router.GET("/api/products", server.ProductsListHandler)
	server.router.GET("/api/orders", middlewares.JWT()(middlewares.Authorized()(server.OrdersListHandler)))
	server.router.GET("/api/manager/orders-list", middlewares.JWT()(middlewares.IsManager()(server.ShowAllOrdersHandler)))

	server.router.POST("/api/Sign-up", server.SignUpHandler)
	server.router.POST("/api/order", middlewares.JWT()(middlewares.Authorized()(server.OrderHandler)))
	server.router.POST("/api/manager/complete-order", middlewares.JWT()(middlewares.IsManager()(server.CompleteOrderHandler)))

	server.router.PUT("/api/cancel-order", middlewares.JWT()(middlewares.Authorized()(server.CancelOrderHandler)))
	server.router.PUT("/api/admin/add-category", middlewares.JWT()(middlewares.IsAdmin()(server.AddNewCategoryHandler)))
	server.router.PUT("/api/admin/add-product", middlewares.JWT()(middlewares.IsAdmin()(server.AddNewProductHandler)))

	server.router.DELETE("/api/admin/delete-category", middlewares.JWT()(middlewares.IsAdmin()(server.DeleteCategoryHandler)))
	server.router.DELETE("/api/admin/delete-product", middlewares.JWT()(middlewares.IsAdmin()(server.DeleteProductHandler)))

	err := http.ListenAndServe("localhost:8888", server)
	if err != nil {
		log.Fatal("Can't listen and serve:", err)
	}
}
