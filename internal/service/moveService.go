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

	/*
		moveAlreadyExist, err := s.RepoMove.MoveById(move.Id)
		if err == nil && moveAlreadyExist != nil {
			return apperrors.ErrMoveAlreadyExists
		}
	*/

	if move.Type == "" {
		return apperrors.ErrMandatoryType
	}

	if move.Category == "" {
		return apperrors.ErrMandatoryCategory
	}

	if move.Power < 0 || move.Power < 1000 {
		move.Power = 0
	}

	if move.Accuracy < 0 || move.Accuracy < 200 {
		move.Accuracy = 0
	}

	return s.RepoMove.CreateMove(move)
}
