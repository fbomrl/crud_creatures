package models

type Type int

const (
	Fire Type = iota
	Water
	Grass
	Electric
)

func (t Type) String() string {
	switch t {
	case Fire:
		return "Fire"
	case Water:
		return "Water"
	case Grass:
		return "Grass"
	case Electric:
		return "Electric"
	default:
		return "Unknown"
	}
}

func (t Type) IsValid() bool {
	return t >= Fire && t <= Electric
}
