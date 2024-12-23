package dto

type UserDTO struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"  validate:"required,min=3,max=32"`
	Email       *string `json:"email" validate:"email"`
	Phone       string  `json:"phone" validate:"required"`
	IsActive    bool    `json:"is_active" validate:"required"`
	CountryCode string  `json:"country_code" validate:"required"`
}

type CreateUserDTO struct {
	Name        string  `json:"name"  validate:"required,min=3,max=32"`
	Phone       string  `json:"email" validate:"required,phone"`
	Email       *string `json:"email" validate:"email"`
	CountryCode string  `json:"country_code" validate:"required"`
}

type UpdateUserDTO struct {
	Name        string  `json:"name"  validate:"omitempty,min=3,max=32"`
	Phone       string  `json:"email" validate:"required,phone"`
	Email       *string `json:"email" validate:"email"`
	CountryCode string  `json:"country_code" validate:"required"`
}
