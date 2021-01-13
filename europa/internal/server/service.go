package server

import (
	"context"
	"log"

	"github.com/hashicorp/go-hclog"
	"gorm.io/gorm"

	"github.com/nessus/europa/internal/domain/entity"
	"github.com/nessus/europa/internal/server/protos"
)

// UserService is
type UserService struct {
	DB  *gorm.DB
	log hclog.Logger
}

// NewUserService is
func NewUserService(db *gorm.DB, logger hclog.Logger) *UserService {
	return &UserService{db, logger}
}

// Authenticate is
func (userService *UserService) Authenticate(ctx context.Context, request *protos.AuthenticationRequest) (*protos.AuthenticationResponse, error) {
	token := request.GetToken()
	db := userService.DB

	u, err := entity.Authenticate(token, db)
	if err != nil {
		res := &protos.AuthenticationResponse{
			User:  nil,
			Error: err.Error(),
		}

		return res, nil
	}

	uPb := &protos.User{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}

	res := &protos.AuthenticationResponse{
		User:  uPb,
		Error: "",
	}

	return res, nil
}

// Register is
func (userService *UserService) Register(ctx context.Context, request *protos.RegisterRequest) (*protos.RegisterResponse, error) {
	db := userService.DB

	u, err := entity.CreateUser(request.GetName(), request.GetEmail(), request.GetPassword(), db)
	if err != nil {
		log.Printf("Error: %s", err.Error())
		return &protos.RegisterResponse{
			Token: "",
			User:  nil,
			Error: err.Error(),
		}, nil
	}

	token, err := entity.LoginUser(request.Email, request.Password)
	if err != nil {
		return &protos.RegisterResponse{
			Token: "",
			User:  nil,
			Error: err.Error(),
		}, nil
	}

	uPb := &protos.User{
		Id:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}

	res := &protos.RegisterResponse{
		Token: token,
		User:  uPb,
		Error: "",
	}

	return res, nil
}

// Login is
func (userService *UserService) Login(ctx context.Context, request *protos.LoginRequest) (*protos.LoginResponse, error) {
	token, err := entity.LoginUser(request.Email, request.Password)

	errStr := ""
	if err != nil {
		errStr = err.Error()
	}

	res := &protos.LoginResponse{
		Token: token,
		Error: errStr,
	}

	return res, nil
}
