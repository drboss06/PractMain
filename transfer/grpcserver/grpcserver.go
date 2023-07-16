package grpc

import (
	"authPract"
	api "authPract/pkg/api"
	"authPract/pkg/service"
	"context"
	"fmt"
	"github.com/spf13/cast"
)

type GRPCServer struct {
	service *service.Service
	api.UnimplementedAuthServer
}

func (g *GRPCServer) GetUserById(ctx context.Context, request *api.GetUserByIdRequest) (*api.GetUserByIdResponse, error) {
	user, err := g.service.GetUserById(cast.ToInt(request.Id))
	if err != nil {
		return &api.GetUserByIdResponse{}, err
	}
	//var userRelay []*api.UserReplay

	//userRelay = append(userRelay, &api.UserReplay{Id: cast.ToInt32(user.Id), Name: user.Name, Username: user.Username, PasswordHash: user.Password})

	return &api.GetUserByIdResponse{Id: cast.ToInt32(user.Id), Name: user.Name, Username: user.Username, PasswordHash: user.Password}, nil
}

func (g *GRPCServer) DeleteUser(ctx context.Context, request *api.DeleteUserRequest) (*api.DeleteUserResponse, error) {
	err := g.service.DeleteUser(cast.ToInt(request.Id))
	if err != nil {
		return &api.DeleteUserResponse{}, err
	}
	return &api.DeleteUserResponse{}, nil
}

func (g *GRPCServer) CreateUser(ctx context.Context, request *api.CreateUserRequest) (*api.CreateUserResponse, error) {
	var input authPract.User
	input.Name = cast.ToString(request.Name)
	input.Username = cast.ToString(request.Username)
	input.Password = cast.ToString(request.Password)
	userId := ctx.Value("userId")
	fmt.Println(userId)

	a, err := g.service.CreateUser(input)
	if err != nil {
		return &api.CreateUserResponse{Id: cast.ToInt32(0)}, err
	}
	return &api.CreateUserResponse{Id: cast.ToInt32(a)}, nil
}

type acceptUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (g *GRPCServer) AcceptUser(ctx context.Context, request *api.AcceptUserRequest) (*api.AcceptUserResponse, error) {
	var input acceptUser
	input.Username = cast.ToString(request.Username)
	input.Password = cast.ToString(request.Password)
	token, err := g.service.GenerateToken(input.Username, input.Password)

	if err != nil {
		return &api.AcceptUserResponse{}, err
	}
	return &api.AcceptUserResponse{Token: token}, nil
}

func NewGrpc(s *service.Service) *GRPCServer {
	return &GRPCServer{service: s}
}
