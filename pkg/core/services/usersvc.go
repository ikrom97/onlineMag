package services

import (
	"database/sql"
	"errors"
	"log"
	"onlineMag/db"
	"onlineMag/hash"
	"onlineMag/models"
)

type UserSvc struct {
	Db *sql.DB
}

func NewUserSvc(Db *sql.DB) *UserSvc {
	if Db == nil {
		log.Println(errors.New("Db can't be nil"))
	}
	return &UserSvc{Db: Db}
}
func (receiver *UserSvc) RegistUser(Db *sql.DB, user models.SignUpBody) (err error) {
	role := "User"
	user.Password, err = hash.HashPassword(user.Password)
	if err != nil {
		log.Println("Can't hash the password:", err)
		return
	}
	_, err = Db.Exec(db.AddNewUser, user.Name, user.Surname, user.Phone, user.Email, role, user.Login, user.Password)
	if err != nil {
		log.Println("Can't add new user:", err)
		return
	}
	return
}
func (receiver *UserSvc) GetUserByLogin(Db *sql.DB, login string) (user models.User, err error) {
	row := Db.QueryRow(db.GetUserByLogin, login)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Phone,
		&user.Email,
		&user.Role,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
func (receiver *UserSvc) CheckHasUser(Db *sql.DB, loginBody models.LoginBody) (user models.User, err error) {
	row := Db.QueryRow(db.GetUserByLogin, loginBody.Login)
	err = row.Scan(
		&user.ID,
		&user.Name,
		&user.Surname,
		&user.Phone,
		&user.Email,
		&user.Role,
		&user.Login,
		&user.Password,
		&user.Remove)
	if err != nil {
		log.Println(err)
		return
	}
	return
}
