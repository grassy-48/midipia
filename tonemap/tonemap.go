package tonemap

const (
	Rest uint8 = 0
	C          = iota + 0x3c
	Cis
	D
	Dis
	E
	F
	Fis
	G
	Gis
	A
	B
	H
	HighC
	HighCis
	HighD
	HighDis
	HighE
)

func Tone(key string) uint8 {
	list := map[string]uint8{
		"a": C,
		"w": Cis,
		"s": D,
		"e": Dis,
		"d": E,
		"f": F,
		"t": Fis,
		"g": G,
		"y": Gis,
		"h": A,
		"u": B,
		"j": H,
		"k": HighC,
		"o": HighCis,
		"l": HighD,
		"p": HighDis,
		";": HighE,
		" ": Rest,
	}
	return list[key]
}
