package repository

import (
	"crud_creatures/internal/models"
	"database/sql"
)

type MovesRepository struct {
	DB *sql.DB
}

func (repo *MovesRepository) CreateMove(moves models.Move) error {
	_, err := repo.DB.Exec(
		"INSERT INTO MOVES (name, type, category, power, accuracy, description) VALUES (?,?,?,?,?,?)",
		&moves.Name,
		&moves.Type,
		&moves.Category,
		&moves.Power,
		&moves.Accuracy,
		&moves.Description,
	)
	return err
}

func (repo *MovesRepository) DeleteMove(moves models.Move) error {
	_, err := repo.DB.Exec(
		"DELETE FROM MOVES WHERE id = ?",
		&moves.Id,
	)
	return err
}

func (repo *MovesRepository) UpdateMove(moves models.Move) error {
	_, err := repo.DB.Exec(
		"UPDATE MOVES SET name = ?, type = ?, category = ?, power = ?, accuracy = ?, description = ? WHERE id = ?",
		&moves.Name,
		&moves.Type,
		&moves.Category,
		&moves.Power,
		&moves.Accuracy,
		&moves.Description,
	)
	return err
}

func (repo *MovesRepository) MoveById(id int) (*models.Move, error) {
	var moves models.Move
	err := repo.DB.QueryRow(
		"SELECT id, name, type, category, power, accuracy, description FROM MOVES WHERE id = ?", id).Scan(
		&moves.Id,
		&moves.Name,
		&moves.Type,
		&moves.Category,
		&moves.Power,
		&moves.Accuracy,
		&moves.Description,
	)
	if err != nil {
		return nil, err
	}
	return &moves, nil
}

func (repo *MovesRepository) FindAllMoves() ([]*models.Move, error) {
	res, err := repo.DB.Query("SELECT id, name, type, category, power, accuracy, description FROM MOVES")
	if err != nil {
		return nil, err
	}

	defer res.Close()

	moves := []*models.Move{}

	for res.Next() {
		var move models.Move
		err := res.Scan(
			&move.Id,
			&move.Name,
			&move.Type,
			&move.Category,
			&move.Power,
			&move.Accuracy,
			&move.Description,
		)
		if err != nil {
			return nil, err
		}
		moves = append(moves, &move)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return moves, nil
}
