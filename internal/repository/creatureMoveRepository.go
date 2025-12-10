package repository

import (
	"database/sql"
)

type CreatureMoveRepository struct {
	DB *sql.DB
}

func (repo *CreatureMoveRepository) AddMoveToCreature(creatureId int, moveId int) error {
	_, err := repo.DB.Exec(
		"INSERT INTO CREATURE_MOVE (id_creature, id_move) VALUES (@p1, @p2)",
		creatureId,
		moveId,
	)

	return err
}

func (repo *CreatureMoveRepository) CreatureExists(id int) (bool, error) {
	var exists int

	err := repo.DB.QueryRow(
		"SELECT TOP 1 1 FROM CREATURE WHERE id = @p1",
		id,
	).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (repo *CreatureMoveRepository) MoveExists(id int) (bool, error) {
	var exists int

	err := repo.DB.QueryRow(
		"SELECT TOP 1 1 FROM MOVES WHERE id = @p1",
		id,
	).Scan(&exists)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (repo *CreatureMoveRepository) CountMovesByCreature(idCreature int) (int, error) {
	var qntdMoves int

	err := repo.DB.QueryRow(`
		SELECT COUNT(*) 
		FROM CREATURE_MOVE
		WHERE id_creature = @p1
	`, idCreature).Scan(&qntdMoves)

	if err != nil {
		return 0, err
	}

	return qntdMoves, nil
}
