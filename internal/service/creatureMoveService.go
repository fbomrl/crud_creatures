package service

import (
	apperrors "crud_creatures/internal/errors"
	repositoryinterface "crud_creatures/internal/repository/interface"
)

type CreatureMoveService struct {
	RepoCreMove repositoryinterface.CreatureMoveInterface
}

func (s *CreatureMoveService) AddMoveToCreatureService(creatureId int, moveId int) error {
	creatureExists, err := s.RepoCreMove.CreatureExists(creatureId)

	if err != nil {
		return err
	}

	if !creatureExists {
		return apperrors.ErrCreatureNoExists
	}

	moveExists, err := s.RepoCreMove.MoveExists(moveId)

	if err != nil {
		return err
	}

	if !moveExists {
		return apperrors.ErrMoveNoExists
	}

	return s.RepoCreMove.AddMoveToCreature(creatureId, moveId)
}
