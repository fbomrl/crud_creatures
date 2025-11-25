package models

type Creature struct {
	Id                 int     `json:"id"`            /*id de criação auto increment*/
	IdGeneral          int     `json:"id_general"`    /*id do general do Creature*/
	IdRegion           int     `json:"id_region"`     /*id daa região do Creature*/
	RegionNumber       int     `json:"region_number"` /*número da região do Creature*/
	Name               string  `json:"name"`
	TypePrimary        string  `json:"type_primary"`
	TypeSecondary      string  `json:"type_secondary"`
	Ability            string  `json:"ability"`
	AbilityDescription string  `json:"ability_description"`
	Species            string  `json:"species"`
	Height             float64 `json:"height"`
	Weight             float64 `json:"weight"`
	Habitat            string  `json:"habitat"`
	Feeding            string  `json:"feeding"`
	Evolution          int     `json:"evolution"`
	Description        string  `json:"description"`
}
