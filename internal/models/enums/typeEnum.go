package enums

type Type int

const (
	Neutral Type = iota
	Fire
	Combat
	Water
	Flying
	Botanical
	Venom
	Electric
	Chthonic
	Mentis
	Mineral
	Frost
	Arthropod
	Dragon
	Spectral
	Shadow
	Metal
	Mythic
	Cosmic
	Sound
	Tribal
	Fungal
	Machine
	Microbial
)

func (t Type) String() string {
	switch t {
	case Neutral:
		return "Neutro"
	case Fire:
		return "Fogo"
	case Combat:
		return "Combate"
	case Water:
		return "Água"
	case Flying:
		return "Voador"
	case Botanical:
		return "Botânico"
	case Venom:
		return "Veneno"
	case Electric:
		return "Elétrico"
	case Chthonic:
		return "Terrestre"
	case Mentis:
		return "Mental"
	case Mineral:
		return "Mineral"
	case Frost:
		return "Gélido"
	case Arthropod:
		return "Artrópode"
	case Dragon:
		return "Dragão"
	case Spectral:
		return "Espectral"
	case Shadow:
		return "Sombrio"
	case Metal:
		return "Metal"
	case Mythic:
		return "Mítico"
	case Cosmic:
		return "Cósmico"
	case Sound:
		return "Som"
	case Tribal:
		return "Tribal"
	case Fungal:
		return "Fungo"
	case Machine:
		return "Máquina"
	case Microbial:
		return "Microbiano"

	default:
		return "Desconhecido"
	}
}

func (t Type) IsValid() bool {
	return t >= Neutral && t <= Microbial
}
