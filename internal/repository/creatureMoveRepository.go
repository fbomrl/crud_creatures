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

func (repo *CreatureMoveRepository) CreatureExists(id int) (bool, error) {
	var creatureExists
	err := r.DB.QueryRow(
		"SELECT 1 FROM CREATURE WHERE id = ? LIMIT 1",
		id,
	).Scan(&creatureExists)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return, err
		}
	}	

	return err, nil
}

func (repo *CreatureMoveRepository) MoveExists(id int) (bool, err) {
	var moveExists
	err := r.DB.QueryRow(
		"SELECT 1 FROM MOVES WHERE id = ? LIMIT 1",
		id,
	)Scan(&moveExists)

	if err != nil {
		if err ==  sql.ErrNoRows {
			return, err
		}
	}

	return err, nil
}