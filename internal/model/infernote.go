package model

type Infernote struct {
	Language Language             `json:"language"`
	Types    map[string]ChordType `json:"types"`
	Chords   map[string]Chord     `json:"chords"`
}
