package dto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Status   string `json:"status" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Role     string `json:"role" validate:"required"`
	Password string `json:"password" validate:"required,min=6"`
	Status   string `json:"status" validate:"required"`
}
