package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/y33550336/link_slider/handler/grpc"
)

type Repository struct {
	*GameRepository

}

var _ grpc.GameStore = (*Repository)(nil)

func NewRepository(db *sqlx.DB) *Repository {
	gameRepo := NewGameRepository(db)

	return &Repository{
		gameRepo,

	}
}
