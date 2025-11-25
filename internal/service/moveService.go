package service

import (
	apperrors "crud_creatures/internal/errors"
	"crud_creatures/internal/models"
	"crud_creatures/internal/repository"
)

type MoveService struct {
	RepoMove repository.MovesRepository
}

func (s *MoveService) CreateMoveService(move models.Move) error {

	if move.Name == "" {
		return apperrors.ErrMandatoryName
	}

	moveAlreadyExist, err := s.RepoMove.MoveById(move.Id)
	if err == nil && moveAlreadyExist != nil {
		return apperrors.ErrMoveAlreadyExists
	}

	return s.RepoMove.CreateMove(move)
}
