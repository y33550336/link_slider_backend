package repository

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/y33550336/link_slider/handler/grpc"
	"github.com/y33550336/link_slider/pb"
)

var _ grpc.GameStore = (*GameRepository)(nil)

type GameRepository struct {
	db *sqlx.DB
}

func NewGameRepository(db *sqlx.DB) *GameRepository {
	return &GameRepository{db: db}
}

func (r *GameRepository) StorePing(ctx context.Context) (*pb.PingResponse, error) {
	// Implement the logic to ping the database or perform any necessary checks.
	// For example, you can execute a simple query to check the connection.
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	err := r.db.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return &pb.PingResponse{}, nil
}

func (r *GameRepository) StoreMovePlayer(ctx context.Context, playerID string, x, y int) (*pb.MovePlayerResponse, error) {
	// Implement the logic to move the player in the database.
	// This is a placeholder implementation; you should replace it with actual database operations.
	if r.db == nil {
		return nil, errors.New("database connection is not initialized")
	}

	// Example: Update player's position in the database.
	query := "UPDATE players SET x = ?, y = ? WHERE id = ?"
	result, err := r.db.ExecContext(ctx, query, x, y, playerID)
	if err != nil {
		return nil, err
	}
	
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return nil, errors.New("no player found with the given ID")
	}


	return &pb.MovePlayerResponse{Message: "Player moved successfully"}, nil
}
