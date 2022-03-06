package tamiat_user

type TamiatUser struct {
	Id       int
	Name     string
	Email    string
	Password string
	RoleId   int
}

type TamiatUserRepository interface {
	Login(tUser TamiatUser) (string, error)
	Create(tUserObj TamiatUser) error
	ReadAll() ([]TamiatUser, error)
	ReadUserByID()
	Update()
	Delete()
}
