package data

import (
	"fmt"

	"github.com/demonyze/infernote/internal/model"
	"github.com/demonyze/infernote/internal/utils"
)

func Init() ([]model.Chord, error) {
	var chords []model.Chord

	data, err := utils.ImportChordsDbGuitar()
	if err != nil {
		return chords, err
	}

	for keyIndex, key := range data.Keys {
		for chordIndex, chordImport := range data.Chords[key] {
			chords = append(chords, model.Chord{
				Id:              fmt.Sprintf("%d%d", keyIndex, chordIndex),
				Name:            fmt.Sprintf("%s %s", chordImport.Key, chordImport.Suffix),
				GuitarPositions: chordImport.Positions,
				Root:            model.Note{Name: chordImport.Key},
				Type:            chordImport.Suffix,
			})
		}
	}

	return chords, nil
}
