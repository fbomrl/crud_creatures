package models

import "crud_creatures/internal/models/enums"

type Move struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Type        enums.Type     `json:"type"`
	Category    enums.Category `json:"category"`
	Power       int            `json:"power"`
	Accuracy    int            `json:"accuracy"`
	Description string         `json:"description"`
}
