package tamiat_user

import "github.com/golang/protobuf/ptypes/timestamp"

type TamiatUser struct {
	Id        int                 `json:"id"`
	CreatedAt timestamp.Timestamp `json:"created_at"`
	UpdatedAt timestamp.Timestamp `json:"updated_at"`
	DeletedAt timestamp.Timestamp `json:"deleted_at"`
	Name      string              `json:"name"`
	Email     string              `json:"email" form:"email" binding:"required,email"`
	Password  string              `json:"password"`
	RoleId    int                 `json:"role_id"`
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
