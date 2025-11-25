package models

import "crud_creatures/internal/models/enums"

type Creature struct {
	Id               int        `json:"id"`            /*id de criação auto increment*/
	IdGeneral        int        `json:"id_general"`    /*id do general do Creature*/
	IdRegion         int        `json:"id_region"`     /*id daa região do Creature*/
	RegionNumber     int        `json:"region_number"` /*número da região do Creature*/
	Name             string     `json:"name"`
	TypePrimary      enums.Type `json:"type_primary"`
	TypeSecondary    enums.Type `json:"type_secondary"`
	AbilityAtack     string     `json:"AbilityAtack"`
	AbilityAtackDesc string     `json:"ability_attack_desc"`
	AbilityDef       string     `json:"AbilityDef"`
	AbilityDefDesc   string     `json:"ability_def_desc"`
	Species          string     `json:"species"`
	Height           float64    `json:"height"`
	Weight           float64    `json:"weight"`
	Habitat          string     `json:"habitat"`
	Feeding          string     `json:"feeding"`
	Evolution        int        `json:"evolution"`
	Description      string     `json:"description"`
	Inspiration      string     `json:"inspiration"`
}
