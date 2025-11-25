package models

type CreatureMove struct {
	Id         int `json:"id"`
	IdCreature int `json:"id_creature"`
	IdMove     int `json:"id_move"`
}
