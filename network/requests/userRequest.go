package requests

type LoginRequest struct {
	Email string `validate:"required,min=3,max=253" json:"email"`
	Password string `validate:"required,min=8,max=16" json:"password"`
}

type RegisterRequest struct {
	Username string `validate:"required,min=3,max=10" json:"username"`
	Email string `validate:"required,min=3,max=253" json:"email"`
	Password string `validate:"required,min=8,max=16" json:"password"`
}


