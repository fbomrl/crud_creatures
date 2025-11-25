package repository

import "crud_creatures/internal/models"

type MovesRepository interface {
	CreateMove(moves models.Move) error
	DeleteMove(moves models.Move) error
	UpdateMove(moves models.Move) error
	MoveById(id int) (*models.Move, error)
	FindAllMoves() ([]*models.Move, error)
}
