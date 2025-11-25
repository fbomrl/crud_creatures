package enums

type Category int

const (
	Physical Category = iota
	Special
	Status
)

func (t Category) String() string {
	switch t {
	case Physical:
		return "Physical"
	case Special:
		return "Special"
	case Status:
		return "Status"
	default:
		return "Unknown"
	}
}

func (t Category) IsValid() bool {
	return t >= Physical && t <= Status
}
