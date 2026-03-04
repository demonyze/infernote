package model

type ChordLang struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Chord struct {
	Id              string                   `json:"id"`
	NameShort       string                   `json:"nameShort"`
	Lang            map[string]ChordLang     `json:"lang"`
	Notes           []Note                   `json:"notes"`
	Root            Note                     `json:"root"`
	GuitarPositions []ChordsDbGuitarPosition `json:"guitarPositions"`
}
