package service

import (
	"github.com/nitkumar91296/banking-auth/domain"
	"github.com/nitkumar91296/banking-auth/dto"
)

type AuthService interface {
	Login(req dto.LoginRequest) (*string, error)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
	// rolePermissions domain.RolePermissions
}

func NewAuthService(repo domain.AuthRepository) AuthService {
	return DefaultAuthService{repo: repo}
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*string, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}

	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (s DefaultAuthService) Verify(urlParams map[string]string) (bool, error) {
	return false, nil
}
