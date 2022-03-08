package tamiat_user

import "github.com/golang/protobuf/ptypes/timestamp"

type TamiatUser struct {
	Id        int
	CreatedAt timestamp.Timestamp
	UpdtedAt  timestamp.Timestamp
	DeletedAt timestamp.Timestamp
	Name      string
	Email     string `json:"email" form:"email" binding:"required,email"`
	Password  string
	RoleId    int
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
