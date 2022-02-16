package models

type ICredentials struct {
	Email    string `json:"email" bindings:"required, email"`
	Password string `json:"password" bindings:"required, gte=6, lte=30"`
}

type IUser struct {
	Name string
	ICredentials
}

type User struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}
