package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A5th model.Chord = model.Chord{
	Id:               model.A5th,
	Name:             data.ChordNames[model.A5th],
	AlternativeNames: data.AlternativeNames[model.A5th],
	Abbreviation:     data.Abbreviations[model.Fith],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.E],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-2-x-x",
				"5-7-7-x-x-x",
				"x-12-14-14-x-x",
				"17-19-19-x-x-x",
			},
		},
	},
}
