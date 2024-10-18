package service

import (
	"github.com/Darkhackit/banking-auth/domain"
	"github.com/Darkhackit/banking-auth/dto"
	"github.com/Darkhackit/banking-auth/errs"
	"golang.org/x/crypto/bcrypt"
)

type LoginService interface {
	Login(login dto.LoginDTO) (*string, *errs.AppError)
	Register(login dto.RegisterDTO) (*dto.RegisterDTO, *errs.AppError)
}

type DefaultLoginService struct {
	repo domain.LoginRepository
}

func (dl DefaultLoginService) Login(login dto.LoginDTO) (*string, *errs.AppError) {
	oriL := domain.Login{
		Username: login.Username,
		Password: login.Password,
	}
	user, err := dl.repo.Login(oriL)
	if err != nil {
		return nil, err
	}
	token, AppError := user.GenerateToken()
	if AppError != nil {
		return nil, errs.NewValidationError("An error occurred" + AppError.Error())
	}
	return token, nil
}
func (dl DefaultLoginService) Register(login dto.RegisterDTO) (*dto.RegisterDTO, *errs.AppError) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(login.Password), bcrypt.DefaultCost)
	password := string(hash)
	oril := domain.Login{
		Username:   login.Username,
		Password:   password,
		AccountID:  login.AccountID,
		CustomerID: login.CustomerID,
		Role:       login.Role,
	}
	user, err := dl.repo.Register(oril)
	if err != nil {
		return nil, err
	}
	oril.Username = user.Username
	return &login, nil
}

func NewLoginService(repo domain.LoginRepository) DefaultLoginService {
	return DefaultLoginService{repo}
}
