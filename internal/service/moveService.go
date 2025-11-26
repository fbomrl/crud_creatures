package service

import (
	apperrors "crud_creatures/internal/errors"
	"crud_creatures/internal/models"
	repositoryinterface "crud_creatures/internal/repository/interface"
	"strings"
)

type MoveService struct {
	RepoMove repositoryinterface.MovesRepositoryInterface
}

func (s *MoveService) CreateMoveService(move models.Move) error {

	if move.Name == "" {
		return apperrors.ErrMandatoryName
	}

	if strings.TrimSpace(move.Name) == "" {
		return apperrors.ErrMandatoryName
	}

	if len(strings.TrimSpace(move.Name)) < 2 {
		return apperrors.ErrNameLenInvalid
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

func (s *MoveService) DeleteMoveService(move models.Move) error {
	moveAlreadyExist, err := s.RepoMove.FindMoveById(move.Id)
	if err != nil || moveAlreadyExist == nil {
		return apperrors.ErrMoveNoExists
	}

	err = s.RepoMove.DeleteMove(move)
	if err != nil {
		return err
	}

	return nil
}

func (s *MoveService) UpdateMoveService(move models.Move) error {

	if move.Name == "" {
		return apperrors.ErrMandatoryName
	}

	if strings.TrimSpace(move.Name) == "" {
		return apperrors.ErrMandatoryName
	}

	if len(strings.TrimSpace(move.Name)) < 2 {
		return apperrors.ErrNameLenInvalid
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

	return s.RepoMove.UpdateMove(move)
}
