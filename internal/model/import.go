package model

type ChordsDbGuitarImport struct {
	Keys   []string                         `json:"keys"`
	Chords map[string][]ChordsDbGuitarChord `json:"chords"`
}

type ChordsDbGuitarChord struct {
	Key       string                   `json:"key"`
	Suffix    string                   `json:"suffix"`
	Positions []ChordsDbGuitarPosition `json:"positions"`
}

type ChordsDbGuitarPosition struct {
	Frets    []int `json:"frets"`
	Fingers  []int `json:"fingers"`
	Barres   []int `json:"barres"`
	Capo     bool  `json:"capo"`
	BaseFret int   `json:"baseFret"`
	Midi     []int `json:"midi"`
}
