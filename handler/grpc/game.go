package grpc

import (
	"context"

	"github.com/y33550336/link_slider/pb"
)

type GameStore interface {
	StorePing(ctx context.Context, message string) (*pb.PingResponse, error)
	StoreMovePlayer(ctx context.Context, playerID string, x, y int) (*pb.MovePlayerResponse, error)
}

type GameHandler struct {
	store GameStore
	pb.UnimplementedGameServiceServer
}

func NewGameHandler(store GameStore) *GameHandler {
	return &GameHandler{store: store}
}

func (h *GameHandler) Ping(ctx context.Context, req *pb.PingRequest) (*pb.PingResponse, error) {
	return h.store.StorePing(ctx, req.Message)
}

func (h *GameHandler) MovePlayer(ctx context.Context, req *pb.MovePlayerRequest) (*pb.MovePlayerResponse, error) {
	return h.store.StoreMovePlayer(ctx, req.PlayerId, int(req.X), int(req.Y))
}

