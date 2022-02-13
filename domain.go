package fiscalizer

type Domain int

const (
	TRADING Domain = 1 + iota
	SERVICES
	GASOIL
	HOTELS
	TAXI
	PARKING
)
