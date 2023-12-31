package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var ADiminished model.Chord = model.Chord{
	Id:               "adiminished",
	Type:             model.Diminished,
	AlternativeNames: data.AlternativeNames[model.Diminished],
	Abbreviation:     data.Abbreviations[model.Diminished],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.DSharp],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-6-7-5-x-x",
				"17-18-19-17-x-x",
				"x-0-1-2-1-x",
				"x-12-13-14-13-x",
			},
		},
	},
}
