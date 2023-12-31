package model

type Chord struct {
	Id               string                `json:"id"`
	Name             ChordNameType         `json:"name"`
	Type             ChordType             `json:"type"`
	AlternativeNames map[Language][]string `json:"alternative_names"`
	Abbreviation     map[Language][]string `json:"abbrevation"`
	Notes            []Note                `json:"notes"`
	Root             Note                  `json:"root"`
	GuitarVariations []GuitarVariation     `json:"guitar_variations"`
}

type GuitarVariation struct {
	Tuning    Tuning   `json:"tuning"`
	Positions []string `json:"positions"`
}
