package service

import (
	"github.com/tamiat/backend/pkg/domain/tamiat_user"
)

type TamiatUserService interface {
	Login(user tamiat_user.TamiatUser) (string, error)
	Create(tUserObj tamiat_user.TamiatUser) (int, error)
	ReadAll() ([]tamiat_user.TamiatUser, error)
	ReadUserByID(id int) (tamiat_user.TamiatUser, error)
	Update(tUserObj tamiat_user.TamiatUser) error
	Delete(id int) error
	GetRoleId(name string) (int, error)
}
type DefaultTamiatUserService struct {
	repo tamiat_user.TamiatUserRepository
}

func (s DefaultTamiatUserService) Login(tUserObj tamiat_user.TamiatUser) (string, error) {
	return s.repo.Login(tUserObj)
}
func (s DefaultTamiatUserService) Create(tUserObj tamiat_user.TamiatUser) (int, error) {
	return s.repo.Create(tUserObj)
}
func (s DefaultTamiatUserService) ReadAll() ([]tamiat_user.TamiatUser, error) {
	return s.repo.ReadAll()
}
func (s DefaultTamiatUserService) ReadUserByID(id int) (tamiat_user.TamiatUser, error) {
	return s.repo.ReadUserByID(id)
}
func (s DefaultTamiatUserService) Update(tUserObj tamiat_user.TamiatUser) error {
	return s.repo.Update(tUserObj)
}
func (s DefaultTamiatUserService) Delete(id int) error {
	return s.repo.Delete(id)
}
func (s DefaultTamiatUserService) GetRoleId(name string) (int, error) {
	return s.repo.GetRoleId(name)
}
