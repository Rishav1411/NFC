package routes

type User struct {
	FirstName string `json:"first_name" validate:"required,min=4,max=25"`
	LastName  string `json:"last_name" validate:"required,min=4,max=25"`
	Phone     string `json:"phone" validate:"required,min=13,max=13,phone"`
}
