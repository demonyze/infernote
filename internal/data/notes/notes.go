package data

import "github.com/demonyze/infernote/internal/model"

type NoteKey = int64

const (
	A NoteKey = iota
	ASharp
	B
	C
	CSharp
	D
	DSharp
	E
	F
	FSharp
	G
	GSharp
)

var Notes = map[NoteKey]model.Note{
	A:      {Name: "A"},
	ASharp: {Name: "A#"},
	B:      {Name: "B"},
	C:      {Name: "C"},
	CSharp: {Name: "C#"},
	D:      {Name: "D"},
	DSharp: {Name: "D#"},
	E:      {Name: "E"},
	F:      {Name: "F"},
	FSharp: {Name: "F#"},
	G:      {Name: "G"},
	GSharp: {Name: "G#"},
}
