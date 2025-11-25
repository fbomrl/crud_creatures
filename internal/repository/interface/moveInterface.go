package repository

import "crud_creatures/internal/models"

type MovesRepositoryInterface interface {
	CreateMove(moves models.Move) error
	DeleteMove(moves models.Move) error
	UpdateMove(moves models.Move) error
	FindMoveById(id int) (*models.Move, error)
	FindAllMoves() ([]*models.Move, error)
}
