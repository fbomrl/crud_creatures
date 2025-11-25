package repository

import "crud_creatures/internal/models"

type CreatureRepository interface {
	CreateCreature(creatures models.Creature) error
	DeleteCreature(creatures models.Creature) error
	UpdateCreature(creatures models.Creature) error
	FindCreatureById(id int) (*models.Creature, error)
	FindAllCreatures() ([]*models.Creature, error)
}
