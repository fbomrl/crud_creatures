package enums

type Category int

const (
	PhysicalAttack Category = iota
	SpecialAttack
	PhysicalDefense
	SpecialDefense
	Status
)

func (t Category) String() string {
	switch t {
	case PhysicalAttack:
		return "Atk Fís"
	case SpecialAttack:
		return "Atk Esp"
	case PhysicalDefense:
		return "Def Fís"
	case SpecialDefense:
		return "Def Esp"
	case Status:
		return "Status"
	default:
		return "Desconhecido"
	}
}

func (t Category) IsValid() bool {
	return t >= PhysicalAttack && t <= Status
}
