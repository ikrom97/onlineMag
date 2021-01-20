package db

const (
	AddNewUser           = `insert into users(name, surname, phone, email, role, login, password) values(($1),($2),($3),($4),($5),($6),($7))`
	AddNewOrder          = `insert into orders(user_id, ordered_date, product_name, address, status) values(($1),($2),($3),($4),($5))`
	AddNewCategory       = `insert into catalog(name) values(($1))`
	AddNewProduct        = `insert into products(category, name, photo, cost, status) values(($1),($2),($3),($4),($5))`
	UpdateCategoryStatus = `update catalog set remove = ($1) where name = ($2)`
	UpdateCategoryByID = `update catalog set remove = false where id = ($1)`
	UpdateProductByName = `update products set status = ($1) where name = ($2)`
	UpdateOrderById = `update orders set updated_date = ($1), status = ($2) where id = ($3)`
	UpdateOrdersStatus   = `update orders set updated_date = ($1), status = ($2) where user_id = ($3) and product_name = ($4)`

	GetOrdersByStatus = `select * from orders where status = ($1)`
	GetCatalogByName = `select * from catalog where name = ($1)`
	GetUserByLogin         = `select * from users where login = ($1)`
	GetCatalogList         = `select * from catalog`
	GetOrdersByUserID      = `select product_name, address, status from orders where user_id = ($1)`
	ShowProductsByCategory = `select name, photo, cost, status from products where category = ($1)`
)
