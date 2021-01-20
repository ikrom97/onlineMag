package app

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"onlineMag/hash"
	"onlineMag/models"
	"onlineMag/token"
)

const (
	MainMenu = `Привет, Я онлайн магазин Mi Tajikistan
1.Каталог
2.Авторизация
3.Регистрация`
	ContentType     = `Content-Type`
	ApplicationJson = `application/json: charset = utf-8`
)

func (server *MainServer) MainHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	_, err := w.Write([]byte(MainMenu))
	if err != nil {
		log.Println("Can't find connection:",err)
		return
	}
}
func (server *MainServer) SignUpHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(ContentType, ApplicationJson)
	var requestBody models.SignUpBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:",err)
			return
		}
		return
	}
	if len(requestBody.Password) < 8 {
		err = json.NewEncoder(w).Encode("Пароль должен содержать более 8 символов")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	err = server.usersvc.RegistUser(server.Db, requestBody)
	if err != nil {
		err := json.NewEncoder(w).Encode("Этот логин уже занят")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		log.Println("Can't regist new user:", err)
		return
	}
	user, err := server.usersvc.GetUserByLogin(server.Db, requestBody.Login)
	if err != nil {
		log.Println("Can't get new user:", err)
		return
	}
	Token := token.CreateToken(user)
	responseToken := models.ResponseToken{
		Description: user.Name,
		Token:       Token,
		Role:        user.Role,
	}
	err = json.NewEncoder(w).Encode(responseToken)
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer)   SignInHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(ContentType, ApplicationJson)
	var requestBody models.LoginBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:",err)
			return
		}
		return
	}
	user, err := server.usersvc.CheckHasUser(server.Db, requestBody)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !hash.CheckPasswordHash(requestBody.Password, user.Password) {
		err := json.NewEncoder(w).Encode("Incorrect password!")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	Token := token.CreateToken(user)
	responseToken := models.ResponseToken{
		Description: user.Name,
		Token:       Token,
		Role:        user.Role,
	}
	err = json.NewEncoder(w).Encode(responseToken)
	if err != nil {
		log.Println("Can't find connection:",err)
		return
	}
}
func (server *MainServer) CatalogHandler(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	w.Header().Set(ContentType, ApplicationJson)
	catalogList, err := models.GetCatalogList(server.Db)
	if err != nil {
		log.Println("Can't get catalog:", err)
		return
	}
	var list []string
	for i := 0; i < len(catalogList); i++ {
		list = append(list, catalogList[i].Name)
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) ProductsListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(ContentType, ApplicationJson)
	var requestBody models.RequestProductsList
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	productsList, err := models.ShowProductsByCategory(server.Db, requestBody)
	if err != nil {
		log.Println("Can't get productsList:", err)
		return
	}
	err = json.NewEncoder(w).Encode(productsList)
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) OrderHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bearerToken := r.Header.Get("Authorization")
	Token := bearerToken[len("Bearer "):]
	claims := token.ParseToken(Token)
	var requestBody models.OrderBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:",err)
			return
		}
		return
	}
	err = models.AddNewOrder(server.Db, claims.ID, requestBody)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.NewEncoder(w).Encode("Ваш запрос принят, мы с вами свяжемся")
	if err != nil {
		log.Println("Can't find connection:",err)
		return
	}
}
func (server *MainServer) OrdersListHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(ContentType, ApplicationJson)
	bearerToken := r.Header.Get("Authorization")
	Token := bearerToken[len("Bearer "):]
	claims := token.ParseToken(Token)
	orderList, err := models.ShowOrders(server.Db, claims.ID)
	err = json.NewEncoder(w).Encode(orderList)
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) AddNewCategoryHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	category, _ := models.CheckHasCategory(server.Db, requestBody)
	if category.Remove == true {
		err := models.AddCategory(server.Db, category)
		if err != nil {
			log.Println("Can't add category:", err)
			return
		}
		err = json.NewEncoder(w).Encode("Category added successfully")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
	}
	err = models.AddNewCategory(server.Db, requestBody)
	if err != nil {
		err := json.NewEncoder(w).Encode("Такая категория уже существует")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		log.Println("Can't add new category:", err)
		return
	}
	err = json.NewEncoder(w).Encode("Category added successfully")
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) AddNewProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody models.AddProductResponse
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	err = models.AddNewProduct(server.Db, requestBody)
	if err != nil {
		log.Println("Can't add new product:", err)
		return
	}
	err = json.NewEncoder(w).Encode("Продукт добавлен")
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) DeleteCategoryHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	err = models.DeleteCategory(server.Db, requestBody)
	if err != nil {
		log.Println("Can't delete category:",err)
		return
	}
	err = json.NewEncoder(w).Encode("Категория успешно удалено")
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) DeleteProductHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody string
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	err = models.DeleteProduct(server.Db, requestBody)
	if err != nil {
		log.Println("Can't delete product:",err)
		return
	}
	err = json.NewEncoder(w).Encode("Продукт успешно удалено")
	if err != nil {
		log.Println("Can't find connection:",err)
		return
	}
}
func (server *MainServer) CompleteOrderHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var requestBody models.OrdersID
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection:", err)
			return
		}
		return
	}
	err = models.CompleteOrder(server.Db, requestBody.OrdersID)
	if err != nil {
		log.Println("Can't complete order:", err)
		return
	}
	err= json.NewEncoder(w).Encode("Заказ успешно выполнен")
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) ShowAllOrdersHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set(ContentType,ApplicationJson)
	status := "Заказано"
	list, err := models.GetOrdersByStatus(server.Db, status)
	if err != nil {
		log.Println("Cant't get orderslist:", err)
		return
	}
	err = json.NewEncoder(w).Encode(list)
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
func (server *MainServer) CancelOrderHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	bearerToken := r.Header.Get("Authorization")
	Token := bearerToken[len("Bearer "):]
	claims := token.ParseToken(Token)
	var requestBody models.OrderBody
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode("json-invalid")
		if err != nil {
			log.Println("Can't find connection")
			return
		}
		return
	}
	err = models.RemoveOrder(server.Db, claims.ID, requestBody)
	if err != nil {
		log.Println("Can't cancel order:",err)
		return
	}
	err= json.NewEncoder(w).Encode("Заказ успешно отменен")
	if err != nil {
		log.Println("Can't find connection:", err)
		return
	}
}
