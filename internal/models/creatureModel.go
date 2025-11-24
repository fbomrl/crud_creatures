package models

type Creature struct {
	Id           int `json:"id"`           /*id de criação auto increment*/
	IdGeneral    int `json:"id_general`    /*id do general do Creature*/
	IdRegion     int `json:"id_region`     /*id da região do Creature*/
	RegionNumber int `json:"region_number` /*número da região do Creature*/
}
