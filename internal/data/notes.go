package data

import "github.com/demonyze/infernote/internal/model"

type NoteKey = string

const (
	A      NoteKey = "A"
	ASharp NoteKey = "A#"
	B      NoteKey = "B"
	C      NoteKey = "C"
	CSharp NoteKey = "C#"
	D      NoteKey = "D"
	DSharp NoteKey = "D#"
	E      NoteKey = "E"
	F      NoteKey = "F"
	FSharp NoteKey = "F#"
	G      NoteKey = "G"
	GSharp NoteKey = "G#"
)

var Notes = map[NoteKey]model.Note{
	A:      {Name: "A"},
	ASharp: {Name: "A#", AlternativeName: "Bb"},
	B:      {Name: "B"},
	C:      {Name: "C"},
	CSharp: {Name: "C#", AlternativeName: "Db"},
	D:      {Name: "D"},
	DSharp: {Name: "D#", AlternativeName: "Eb"},
	E:      {Name: "E"},
	F:      {Name: "F"},
	FSharp: {Name: "F#", AlternativeName: "Gb"},
	G:      {Name: "G"},
	GSharp: {Name: "G#", AlternativeName: "Ab"},
}
