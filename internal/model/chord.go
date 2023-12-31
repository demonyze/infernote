package model

type Chord struct {
	Id               string               `json:"id"`
	Name             ChordNameType        `json:"name"`
	AlternativeNames AlternativeNamesType `json:"alternative_names"`
	Abbreviation     AbbreviationType     `json:"abbrevation"`
	Notes            []Note               `json:"notes"`
	Root             Note                 `json:"root"`
	GuitarVariations []GuitarVariation    `json:"guitar_variations"`
}

type GuitarVariation struct {
	Tuning    Tuning   `json:"tuning"`
	Positions []string `json:"positions"`
}
