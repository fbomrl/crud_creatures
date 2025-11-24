package interface

type ICreaturesRepository interface {
	CreateCreatures(fakemon *models.Creature) error
	GetCreaturesById(id string) (*models.Creature, error)
	GetAllCreatures() ([]*models.Creature, error)
	UpdateCreatures *models.Creature) error
	DeleteCreatures(id string) error
}