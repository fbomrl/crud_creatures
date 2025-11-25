package repository

import (
	"crud_creatures/internal/models"
	"database/sql"
)

type CreatureRepository struct {
	DB *sql.DB
}

func (repo *CreatureRepository) CreateCreature(creature models.Creature) error {
	_, err := repo.DB.Exec(
		"INSERT INTO CREATURE (id_general, id_region, region_number, name, type_primary, type_secondary, ability, ability_description, species, height, weight, habitat, feeding, evolution, description) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)",
		&creature.IdGeneral,
		&creature.IdRegion,
		&creature.RegionNumber,
		&creature.Name,
		&creature.TypePrimary,
		&creature.TypeSecondary,
		&creature.Ability,
		&creature.AbilityDescription,
		&creature.Species,
		&creature.Height,
		&creature.Weight,
		&creature.Habitat,
		&creature.Feeding,
		&creature.Evolution,
		&creature.Description,
	)
	return err
}

func (repo *CreatureRepository) DeleteCreature(creature *models.Creature) error {
	_, err := repo.DB.Exec(
		"DELETE FROM CREATURE  WHERE ID = ?",
		&creature.Id,
	)
	return err
}

func (repo *CreatureRepository) UpdateCreature(creature *models.Creature) error {
	_, err := repo.DB.Exec(
		"UPDATE CREATURE SET id_general = ?, id_region = ?, region_number = ?, name = ?, type_primary = ?, type_secondary = ?, ability = ?, ability_description = ?, species = ?, height = ?, weight = ?, habitat = ?, feeding = ?, evolution = ?, description = ? WHERE id = ?",
		&creature.IdGeneral,
		&creature.IdRegion,
		&creature.RegionNumber,
		&creature.Name,
		&creature.TypePrimary,
		&creature.TypeSecondary,
		&creature.Ability,
		&creature.AbilityDescription,
		&creature.Species,
		&creature.Height,
		&creature.Weight,
		&creature.Habitat,
		&creature.Feeding,
		&creature.Evolution,
		&creature.Description,
	)
	return err
}

func (repo *CreatureRepository) FindCreatureById(id int) (*models.Creature, error) {
	var creature models.Creature
	err := repo.DB.QueryRow(
		"SELECT id, id_general, id_region, region_number, name, type_primary, type_secondary, ability, ability_description, species, height, weight, habitat, feeding, evolution, description FROM CREATURE WHERE id = ?", id).Scan(
		&creature.Id,
		&creature.IdGeneral,
		&creature.IdRegion,
		&creature.RegionNumber,
		&creature.Name,
		&creature.TypePrimary,
		&creature.TypeSecondary,
		&creature.Ability,
		&creature.AbilityDescription,
		&creature.Species,
		&creature.Height,
		&creature.Weight,
		&creature.Habitat,
		&creature.Feeding,
		&creature.Evolution,
		&creature.Description,
	)
	if err != nil {
		return nil, err
	}
	return &creature, nil
}

func (repo *CreatureRepository) FindCreatureByName(name string) (*models.Creature, error) {
	var creatureName models.Creature
	err := repo.DB.QueryRow(
		"SELECT name WHERE name = ?", name).Scan(
		&creatureName.Name,
	)

	if err != nil {
		return nil, err
	}

	return &creatureName, nil
}

func (repo *CreatureRepository) FindAllCreatures() ([]*models.Creature, error) {
	res, err := repo.DB.Query("SELECT id, id_general, id_region, region_number, name, type_primary, type_secondary, ability, ability_description, species, height, weight, habitat, feeding, evolution, description FROM CREATURE")
	if err != nil {
		return nil, err
	}

	defer res.Close()

	creatures := []*models.Creature{}

	for res.Next() {
		var creature models.Creature
		err := res.Scan(
			&creature.Id,
			&creature.IdGeneral,
			&creature.IdRegion,
			&creature.RegionNumber,
			&creature.Name,
			&creature.TypePrimary,
			&creature.TypeSecondary,
			&creature.Ability,
			&creature.AbilityDescription,
			&creature.Species,
			&creature.Height,
			&creature.Weight,
			&creature.Habitat,
			&creature.Feeding,
			&creature.Evolution,
			&creature.Description,
		)
		if err != nil {
			return nil, err
		}

		creatures = append(creatures, &creature)
	}

	if err = res.Err(); err != nil {
		return nil, err
	}

	return creatures, nil
}
