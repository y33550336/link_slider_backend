package handler

import (
	game "github.com/y33550336/link_slider/handler/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	testgen "github.com/y33550336/link_slider/pb"
	"github.com/y33550336/link_slider/repository"
)

func newGRPCServer(repo *repository.Repository) *grpc.Server {
	grpcServer := grpc.NewServer()
	
	gameHandler := game.NewGameHandler(repo)
	testgen.RegisterGameServiceServer(grpcServer, gameHandler)

	reflection.Register(grpcServer)

	return grpcServer
}
