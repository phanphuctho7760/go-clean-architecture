package entities

type User struct {
	Id       string `json:"id"`
	UserName string `json:"user_name" validate:"required,lowercase"`
}
