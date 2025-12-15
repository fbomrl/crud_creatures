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
		"INSERT INTO CREATURE (id_general, id_region, region_number, name, type_primary, type_secondary, ability_attack, ability_attack_desc, ability_def, ability_def_desc, species, height, weight, habitat, feeding, evolution, description, inspiration) VALUES (@p1,@p2,@p3,@p4,@p5,@p6,@p7,@p8,@p9,@p10,@p11,@p12,@p13,@p14,@p15,@p16,@p17,@p18)",
		creature.IdGeneral,
		creature.IdRegion,
		creature.RegionNumber,
		creature.Name,
		creature.TypePrimary,
		creature.TypeSecondary,
		creature.AbilityAtack,
		creature.AbilityAtackDesc,
		creature.AbilityDef,
		creature.AbilityDefDesc,
		creature.Species,
		creature.Height,
		creature.Weight,
		creature.Habitat,
		creature.Feeding,
		creature.Evolution,
		creature.Description,
		creature.Inspiration,
	)
	return err
}

func (repo *CreatureRepository) DeleteCreature(creature models.Creature) error {
	_, err := repo.DB.Exec(
		"DELETE FROM CREATURE WHERE id = @p1",
		creature.Id,
	)
	return err
}

func (repo *CreatureRepository) UpdateCreature(creature models.Creature) error {
	_, err := repo.DB.Exec(
		`UPDATE CREATURE SET 
			id_general = @p1, 
			id_region = @p2, 
			region_number = @p3, 
			name = @p4, 
			type_primary = @p5, 
			type_secondary = @p6, 
			ability_attack = @p7, 
			ability_attack_desc = @p8, 
			ability_def = @p9, 
			ability_def_desc = @p10, 
			species = @p11, 
			height = @p12, 
			weight = @p13, 
			habitat = @p14, 
			feeding = @p15, 
			evolution = @p16, 
			description = @p17, 
			inspiration = @p18 
		WHERE id = @p19`,
		creature.IdGeneral,
		creature.IdRegion,
		creature.RegionNumber,
		creature.Name,
		creature.TypePrimary,
		creature.TypeSecondary,
		creature.AbilityAtack,
		creature.AbilityAtackDesc,
		creature.AbilityDef,
		creature.AbilityDefDesc,
		creature.Species,
		creature.Height,
		creature.Weight,
		creature.Habitat,
		creature.Feeding,
		creature.Evolution,
		creature.Description,
		creature.Inspiration,
		creature.Id,
	)
	return err
}

func (repo *CreatureRepository) FindCreatureById(id int) (*models.Creature, error) {
	var creature models.Creature
	err := repo.DB.QueryRow(
		`SELECT 
			id, id_general, id_region, region_number, name, 
			type_primary, type_secondary, ability_attack, ability_attack_desc, 
			ability_def, ability_def_desc, species, height, weight, habitat, 
			feeding, evolution, description, inspiration 
		FROM CREATURE WHERE id = @p1`, id).Scan(
		&creature.Id,
		&creature.IdGeneral,
		&creature.IdRegion,
		&creature.RegionNumber,
		&creature.Name,
		&creature.TypePrimary,
		&creature.TypeSecondary,
		&creature.AbilityAtack,
		&creature.AbilityAtackDesc,
		&creature.AbilityDef,
		&creature.AbilityDefDesc,
		&creature.Species,
		&creature.Height,
		&creature.Weight,
		&creature.Habitat,
		&creature.Feeding,
		&creature.Evolution,
		&creature.Description,
		&creature.Inspiration,
	)
	if err != nil {
		return nil, err
	}
	return &creature, nil
}

func (repo *CreatureRepository) FindCreatureByName(name string) (*models.Creature, error) {
	var creatureName models.Creature
	err := repo.DB.QueryRow(
		"SELECT name FROM CREATURE WHERE name = @p1", name).Scan(
		&creatureName.Name,
	)

	if err != nil {
		return nil, err
	}

	return &creatureName, nil
}

func (repo *CreatureRepository) FindAllCreatures() ([]*models.Creature, error) {
	res, err := repo.DB.Query(`
		SELECT 
			id, id_general, id_region, region_number, name, 
			type_primary, type_secondary, ability_attack, ability_attack_desc, 
			ability_def, ability_def_desc, species, height, weight, habitat, 
			feeding, evolution, description, inspiration 
		FROM CREATURE`)
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
			&creature.AbilityAtack,
			&creature.AbilityAtackDesc,
			&creature.AbilityDef,
			&creature.AbilityDefDesc,
			&creature.Species,
			&creature.Height,
			&creature.Weight,
			&creature.Habitat,
			&creature.Feeding,
			&creature.Evolution,
			&creature.Description,
			&creature.Inspiration,
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
