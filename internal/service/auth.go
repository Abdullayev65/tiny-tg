package service

import (
	"tiny-tg/internal/dtos"
	"tiny-tg/internal/models"
	"tiny-tg/internal/pkg/app_errors"
	"tiny-tg/internal/pkg/jwt_manager"
	"tiny-tg/internal/repository"
)

type Auth struct {
	Repo       *repository.Repo
	JwtManager *jwt_manager.JwtManager
}

func (s *Auth) Register(data *models.User) (*dtos.AuthRes, error) {
	user, err := s.register(data)
	if err != nil {
		return nil, err
	}

	return s.authRes(user)
}

func (s *Auth) Login(data *dtos.Login) (*dtos.AuthRes, error) {
	if data.Username == "" || data.Password == "" {
		return nil, app_errors.AuthLoginDataRequired
	}

	user, err := s.Repo.Users.GetByUsername(data.Username)
	if err != nil {
		return nil, app_errors.AuthUserNotFound
	}

	if user.Password != data.Password {
		return nil, app_errors.AuthInvalidPassword
	}

	return s.authRes(user)
}

func (s *Auth) register(data *models.User) (*models.User, error) {
	if data.Username == "" {
		return nil, app_errors.AuthEmptyUsernameOrEmail
	}

	if data.Password == "" {
		return nil, app_errors.AuthEmptyPassword
	}

	userByUsername, _ := s.Repo.Users.GetByUsername(data.Username)
	if userByUsername != nil && userByUsername.ID > 0 {
		return nil, app_errors.AuthDuplicateUsername
	}

	user, err := s.Repo.Users.Create(data)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Auth) authRes(user *models.User) (*dtos.AuthRes, error) {

	token, err := s.JwtManager.Generate(user.ID)
	if err != nil {
		return nil, err
	}

	res := &dtos.AuthRes{
		User:  user,
		Token: token,
	}

	return res, nil
}
