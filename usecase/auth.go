package usecase

import (
	"context"

	"github.com/aldysp34/sm_padang/dto"
	"github.com/aldysp34/sm_padang/repository"
)

type AuthUsecase struct {
	userRepo *repository.UserRepository
}

func NewAuthUsecase(user *repository.UserRepository) *AuthUsecase {
	return &AuthUsecase{
		userRepo: user,
	}
}

func (au *AuthUsecase) Login(ctx context.Context, req dto.ReqUser) (*dto.LoginResponse, error) {
	var response dto.LoginResponse
	var token string

	user, err := au.userRepo.GetUserByUsername(ctx, req)
	if err != nil {
		return nil, err
	}

	token, _ = dto.GenerateAccessToken(dto.JWTClaims{
		Role: user.RoleID,
		ID:   user.ID,
	})

	response.Token = token
	response.Nama = user.Nama
	response.Username = user.Username
	response.Role = user.Role.RoleName
	response.RoleID = user.RoleID

	return &response, nil
}
