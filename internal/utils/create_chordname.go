package utils

import (
	"fmt"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
)

func GetChordName(chord model.Chord, lang model.Language) string {
	root := chord.Root.Name
	chordType := chord.Type

	if chordNames, ok := data.ChordNames[chordType]; ok {
		if name, langOk := chordNames[lang]; langOk {
			return fmt.Sprintf("%s %s", root, name)
		}
	}

	return ""
}
