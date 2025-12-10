package repositoryinterface

type CreatureMoveInterface interface {
	AddMoveToCreature(creatureId int, moveId int) error
	CreatureExists(id int) (bool, error)
	MoveExists(id int) (bool, error)
	CountMovesByCreature(idCreature int) (int, error)
}
