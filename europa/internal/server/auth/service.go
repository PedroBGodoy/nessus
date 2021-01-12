package auth

import (
	"context"
	"log"

	"github.com/nessus/europa/internal/models"
	"github.com/nessus/europa/internal/session"
)

// UserService is
type UserService struct {
	ctx context.Context
}

// Authenticate is
func (userService *UserService) Authenticate(ctx context.Context, request *AuthenticationRequest) (*AuthenticationResponse, error) {
	token := request.GetToken()

	db := session.GetDatabase(userService.ctx)

	u, err := models.Authenticate(token, db)
	if err != nil {
		res := &AuthenticationResponse{
			User:  nil,
			Error: err.Error(),
		}

		return res, nil
	}

	uPb := &User{
		Id:    u.UserID,
		Name:  u.Name,
		Email: u.Email,
	}

	res := &AuthenticationResponse{
		User:  uPb,
		Error: "",
	}

	return res, nil
}

// Register is
func (userService *UserService) Register(ctx context.Context, request *RegisterRequest) (*RegisterResponse, error) {
	db := session.GetDatabase(userService.ctx)

	u, err := models.CreateUser(request.GetName(), request.GetEmail(), request.GetPassword(), db)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return &RegisterResponse{
			Token: "",
			User:  nil,
			Error: err.Error(),
		}, nil
	}

	token, err := models.LoginUser(request.Email, request.Password)
	if err != nil {
		return &RegisterResponse{
			Token: "",
			User:  nil,
			Error: err.Error(),
		}, nil
	}

	uPb := &User{
		Id:    u.UserID,
		Name:  u.Name,
		Email: u.Email,
	}

	res := &RegisterResponse{
		Token: token,
		User:  uPb,
		Error: "",
	}

	return res, nil
}

// Login is
func (userService *UserService) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	token, err := models.LoginUser(request.Email, request.Password)

	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	res := &LoginResponse{
		Token: token,
		Error: errStr,
	}

	return res, nil
}
