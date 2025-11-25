package service

import (
	"crud_creatures/internal/models"
	"crud_creatures/internal/repository"
)

type MoveService struct {
	RepoMove repository.MovesRepository
}

func (s *MoveService) CreateMoveService(move models.Move) error {

	return s.RepoMove.CreateMove(move)
}
