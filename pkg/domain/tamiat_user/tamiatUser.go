package tamiat_user

import "github.com/golang/protobuf/ptypes/timestamp"

type TamiatUser struct {
	Id        int
	CreatedAt timestamp.Timestamp
	UpdtedAt  timestamp.Timestamp
	DeletedAt timestamp.Timestamp
	Name      string
	Email     string
	Password  string
	RoleId    int
}

type TamiatUserRepository interface {
	Login(tUser TamiatUser) (string, error)
	Create(tUserObj TamiatUser) error
	ReadAll() ([]TamiatUser, error)
	ReadUserByID(id int) error
	Update(tUserObj TamiatUser) error
	Delete(id int) error
}
