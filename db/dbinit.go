package db

import "database/sql"

func DatabaseInit(Db *sql.DB) (err error) {
	DDLs := []string{CreateUsersTable, CreateCatalogsTable, CreateProductsTable, CreateOrdersTable}
	for _, ddl := range DDLs {
		_, err = Db.Exec(ddl)
		if err != nil {
			return err
		}
	}
	return
}
