package adder

import (
	"authPract"
	api "authPract/pkg/api"
	"authPract/pkg/service"
	"context"
	"github.com/spf13/cast"
)

type GRPCServer struct {
	service *service.Service
	api.UnimplementedAdderServer
}

func (g *GRPCServer) AddUserToTeam(ctx context.Context, request *api.AddUserToTeamRequest) (*api.AddUserToTeamResponse, error) {
	a, err := g.service.AddUserToTeam(cast.ToInt(ctx.Value("userId")), cast.ToInt(request.TeamId))
	if err != nil {
		return &api.AddUserToTeamResponse{Id: cast.ToInt32(0)}, err
	}
	return &api.AddUserToTeamResponse{Id: cast.ToInt32(a)}, nil
}

func (g *GRPCServer) SendMailToUser(ctx context.Context, request *api.UserEmailRequest) (*api.UserEmailResponse, error) {
	err := g.service.SendMailToUser(request.Email)
	if err != nil {
		return &api.UserEmailResponse{Ansver: cast.ToString(err.Error())}, err
	}
	return &api.UserEmailResponse{Ansver: "ok"}, nil
}

func (g *GRPCServer) ParseToken(ctx context.Context, request *api.AccessTokenRequest) (*api.AccessTokenResponse, error) {
	userId, err := g.service.ParseToken(request.Token)
	_ = context.WithValue(ctx, "userId", userId)
	a := ctx.Value("userId")

	if err != nil {
		return &api.AccessTokenResponse{TokenId: cast.ToInt32(0)}, err
	}
	return &api.AccessTokenResponse{TokenId: cast.ToInt32(a)}, nil
}

func (g *GRPCServer) CreateTeam(ctx context.Context, request *api.CreateTeamRequest) (*api.CreateTeamResponse, error) {
	var input authPract.Team

	input.Name = request.Name
	input.Description = request.Description

	a, err := g.service.CreateTeam(cast.ToInt(ctx.Value("userId")), input)
	if err != nil {
		return &api.CreateTeamResponse{Id: cast.ToInt32(1)}, nil
	}

	return &api.CreateTeamResponse{Id: cast.ToInt32(a)}, nil
}

func (g *GRPCServer) mustEmbedUnimplementedAdderServer() {
	//TODO implement me
	panic("implement me")
}

func NewGrpc(services *service.Service) *GRPCServer {
	return &GRPCServer{service: services}
}

//func (g *GRPCServer) CreateTeam(ctx *context.Context, in *api.CreateTeamRequest) (*api.CreateTeamResponse, error) {
//	var input authPract.Team
//	input.Name = in.Team.Name
//	input.Description = in.Team.Description
//
//	a, _ := g.service.CreateTeam(cast.ToInt(in.Id), input)
//
//	return &api.CreateTeamResponse{Id: cast.ToInt32(a)}, nil
//}
