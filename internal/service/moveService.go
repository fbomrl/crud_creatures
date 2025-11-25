package service

import (
	apperrors "crud_creatures/internal/errors"
	"crud_creatures/internal/models"
	repositoryinterface "crud_creatures/internal/repository/interface"
)

type MoveService struct {
	RepoMove repositoryinterface.MovesRepositoryInterface
}

func (s *MoveService) CreateMoveService(move models.Move) error {

	if move.Name == "" {
		return apperrors.ErrMandatoryName
	}

	moveAlreadyExist, err := s.RepoMove.FindMoveByName(move.Name)
	if err == nil && moveAlreadyExist != nil {
		return apperrors.ErrMoveAlreadyExists
	}

	if !move.Type.IsValid() {
		return apperrors.ErrMandatoryType
	}

	if !move.Category.IsValid() {
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
