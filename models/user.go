package models

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    int64  `json:"phone"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Remove   bool   `json:"remove"`
}
type SignUpBody struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Phone    int64  `json:"phone"`
	Email    string `json:"email"`
	Login    string `json:"login"`
	Password string `json:"password"`
}
type ResponseToken struct {
	Description string `json:"description"`
	Token       string `json:"token"`
	Role        string `json:"role"`
}
type LoginBody struct {
	Login    string `json:"login, omitempty"`
	Password string `json:"password, omitempty"`
}
