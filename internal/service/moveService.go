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

	moveAllreadyExists, err := s.RepoMove.FindMoveById(move.Id)
	if err != nil || moveAllreadyExists == nil {
		return apperrors.ErrMoveNoExists
	}

	moveSameName, err := s.RepoMove.FindMoveByName(move.Name)
	if err != nil || moveSameName == nil {
		return apperrors.ErrMoveSameName
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

	moveAllreadyExists, err := s.RepoMove.FindMoveById(move.Id)
	if err != nil || moveAllreadyExists == nil {
		return apperrors.ErrMoveNoExists
	}

	moveSameName, err := s.RepoMove.FindMoveByName(move.Name)
	if err != nil || moveSameName == nil {
		return apperrors.ErrMoveSameName
	}

	return s.RepoMove.UpdateMove(move)
}

func (s *MoveService) FindMoveByIdService(id int) (*models.Move, error) {
	moveId, err := s.RepoMove.FindMoveById(id)
	if err != nil || moveId == nil {
		return nil, apperrors.ErrMoveNoExists
	}

	return moveId, nil
}

func (s *MoveService) FindMoveByNameService(name string) (*models.Move, error) {
	moveName, err := s.RepoMove.FindMoveByName(name)
	if err != nil || moveName == nil {
		return nil, apperrors.ErrMoveNoExists
	}

	return moveName, nil
}

func (s *MoveService) FindAllMovesService() ([]*models.Move, error) {
	allMoves, err := s.RepoMove.FindAllMoves()
	if err != nil || allMoves == nil {
		return nil, apperrors.ErrMoveNoExists
	}

	return allMoves, nil
}
