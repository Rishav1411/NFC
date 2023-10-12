package routes

type User struct {
	Name  string `json:"name" validate:"required,min=4,max=50"`
	Reg   string `json:"reg" validate:"required,min=9,max=9,reg"`
	Phone string `json:"phone" validate:"required,min=13,max=13,phone"`
}

type OTP struct {
	Phone string `json:"phone" validate:"required,min=13,max=13,phone"`
	Otp   string `json:"otp" validate:"required,min=4,max=4,otp"`
}
