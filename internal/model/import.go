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

// Chord Rocks
type ChordRocksGuitarChord struct {
	InitialSuffixName string `json:"internal_suffix_name"`
	Notes             string `json:"notes"`
	Degrees           string `json:"degrees"`
	Abbrevation       string `json:"abbrevation"`
	AlternativeNames  string `json:"alternative_names"`
}
type ChordRocksGuitarChordMap = map[string]ChordRocksGuitarChord
type ChordRocksGuitarImport = map[string]ChordRocksGuitarChordMap
