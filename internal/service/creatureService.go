package service

import (
	apperrors "crud_creatures/internal/errors"
	"crud_creatures/internal/models"
	"crud_creatures/internal/models/enums"
	repositoryinterface "crud_creatures/internal/repository/interface"
	"strings"
)

type CreatureService struct {
	RepoCreature repositoryinterface.CreatureRepositoryInterface
}

func (s *CreatureService) CreateCreatureService(creature models.Creature) error {
	if creature.IdGeneral == 0 || creature.IdRegion == 0 {
		return apperrors.ErrMandatoryIdGeneralOrRegion
	}

	if creature.IdGeneral < 0 || creature.IdRegion < 0 {
		return apperrors.ErrNegativeId
	}

	//
	//REGION NUMBER
	//

	if creature.Name == "" {
		return apperrors.ErrMandatoryName
	}

	if strings.TrimSpace(creature.Name) == "" {
		return apperrors.ErrMandatoryName
	}

	if len(creature.Name) < 2 {
		return apperrors.ErrNameLenInvalid
	}

	creatureSameName, err := s.RepoCreature.FindCreatureByName(creature.Name)
	if err == nil || creatureSameName != nil {
		return apperrors.ErrCreatureSameName
	}

	if !creature.TypePrimary.IsValid() {
		return apperrors.ErrMandatoryType
	}

	if creature.TypeSecondary == 0 {
		creature.TypeSecondary = enums.None
	}

	if creature.Height < 0 {
		return apperrors.ErrNegativeHeightProhibited
	}

	if creature.Weight < 0 {
		return apperrors.ErrNegativeWeightProhibited
	}

	return s.RepoCreature.CreateCreature(creature)
}

func (s *CreatureService) DeleteCreatureService(creature models.Creature) error {
	creatureExists, err := s.RepoCreature.FindCreatureById(creature.Id)
	if err != nil || creatureExists == nil {
		return apperrors.ErrCreatureNoExists
	}

	err = s.RepoCreature.DeleteCreature(creature)
	if err != nil {
		return err
	}
	return nil
}

func (s *CreatureService) UpdateCreatureService(creature models.Creature) error {

	if creature.IdGeneral == 0 || creature.IdRegion == 0 {
		return apperrors.ErrMandatoryIdGeneralOrRegion
	}

	if creature.IdGeneral < 0 || creature.IdRegion < 0 {
		return apperrors.ErrNegativeId
	}

	//
	//REGION NUMBER
	//

	if creature.Name == "" {
		return apperrors.ErrMandatoryName
	}

	if strings.TrimSpace(creature.Name) == "" {
		return apperrors.ErrMandatoryName
	}

	if len(creature.Name) < 2 {
		return apperrors.ErrNameLenInvalid
	}

	creatureSameName, err := s.RepoCreature.FindCreatureByName(creature.Name)
	if err == nil && creatureSameName != nil && creatureSameName.Id != creature.Id {
		return apperrors.ErrCreatureSameName
	}

	if !creature.TypePrimary.IsValid() {
		return apperrors.ErrMandatoryType
	}

	if creature.TypeSecondary == 0 {
		creature.TypeSecondary = enums.None
	}

	if creature.Height < 0 {
		return apperrors.ErrNegativeHeightProhibited
	}

	if creature.Weight < 0 {
		return apperrors.ErrNegativeWeightProhibited
	}

	return s.RepoCreature.UpdateCreature(creature)
}

func (s *CreatureService) FindCreatureByIdService(id int) (*models.Creature, error) {
	creatureId, err := s.RepoCreature.FindCreatureById(id)

	if err != nil || creatureId == nil {
		return nil, apperrors.ErrCreatureNoExists
	}

	return creatureId, nil
}

func (s *CreatureService) FindCreatureByNameService(name string) (*models.Creature, error) {
	creatureName, err := s.RepoCreature.FindCreatureByName(name)

	if err != nil || creatureName == nil {
		return nil, apperrors.ErrCreatureNoExists
	}

	return creatureName, nil
}

func (s *CreatureService) FindAllCreaturesService() ([]*models.Creature, error) {
	allCreatures, err := s.RepoCreature.FindAllCreatures()

	if err != nil || allCreatures == nil {
		return nil, apperrors.ErrCreatureListEmpty
	}

	return allCreatures, nil
}
