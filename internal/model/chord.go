package model

type Chord struct {
	Id               string                   `json:"id"`
	Name             string                   `json:"name"`
	Type             ChordType                `json:"type"`
	AlternativeNames map[Language][]string    `json:"alternativeNames"`
	Abbreviation     map[Language][]string    `json:"abbrevation"`
	Notes            []Note                   `json:"notes"`
	Root             Note                     `json:"root"`
	GuitarPositions  []ChordsDbGuitarPosition `json:"guitarPositions"`
}
