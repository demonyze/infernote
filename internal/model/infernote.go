package model

type Infernote struct {
	Languages []string             `json:"languages"`
	Types     map[string]ChordType `json:"types"`
	Chords    map[string]Chord     `json:"chords"`
}
