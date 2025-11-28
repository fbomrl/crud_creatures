package repository

import (
	"database/sql"
)

type CreatureMoveRepository struct {
	DB *sql.DB
}

func (repo *CreatureMoveRepository) AddMoveToCreature(creatureId int, moveId int) error {
	_, err := repo.DB.Exec(
		"INSERT INTO CREATURE_MOVE (id_creature, id_move) VALUES (?,?)",
		creatureId,
		moveId,
	)

	return err
}
