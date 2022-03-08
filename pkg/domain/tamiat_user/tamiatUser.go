package tamiat_user

import (
	"time"
)

type TamiatUser struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email" form:"email" binding:"required,email"`
	Password  string    `json:"password"`
	RoleId    int       `json:"role_id"`
}

type TamiatUserRepository interface {
	Login(tUser TamiatUser) (string, error)
	Create(tUserObj TamiatUser) (int, error)
	ReadAll() ([]TamiatUser, error)
	ReadUserByID(id int) (TamiatUser, error)
	Update(tUserObj TamiatUser) error
	ResetPassword(tUserObj TamiatUser) error
	Delete(id int) error
	GetRoleId(name string) (int, error)
}
