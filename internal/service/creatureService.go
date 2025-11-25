package service

import (
	apperrors "crud_creatures/internal/errors"
	"crud_creatures/internal/models"
	repositoryinterface "crud_creatures/internal/repository/interface"
)

type CreatureService struct {
	RepoCreature repositoryinterface.CreatureRepositoryInterface
}

func (s *CreatureService) CreateCreatureService(creature models.Creature) error {

	if creature.Name == "" {
		return apperrors.ErrMandatoryName
	}

	return s.RepoCreature.CreateCreature(creature)
}
