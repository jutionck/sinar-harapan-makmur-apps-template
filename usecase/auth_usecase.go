package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type AuthenticationUseCase interface {
	Login(username string, password string) (*model.UserCredential, error)
}

type authenticationUseCase struct {
	repo repository.UserRepository
}

func (a *authenticationUseCase) Login(username string, password string) (*model.UserCredential, error) {
	user, err := a.repo.GetByUsernamePassword(username, password)
	if err != nil {
		return nil, fmt.Errorf("User with username: %s not found", username)
	}
	return user, nil
}

func NewAuthenticationUseCase(repo repository.UserRepository) AuthenticationUseCase {
	return &authenticationUseCase{repo: repo}
}
