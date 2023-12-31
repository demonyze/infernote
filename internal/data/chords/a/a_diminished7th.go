package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var ADiminished7th model.Chord = model.Chord{
	Id:               "adiminished7th",
	Type:             model.Diminished7th,
	AlternativeNames: data.AlternativeNames[model.Diminished7th],
	Abbreviation:     data.Abbreviations[model.Diminished7th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.DSharp],
		notes.Notes[notes.FSharp],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-1-2-1-2",
				"5-6-4-5-4-x",
				"5-6-7-5-7-5",
				"x-x-7-8-7-8",
				"x-12-10-11-10-11",
				"x-12-13-11-13-11",
				"17-18-16-17-16-x",
				"17-18-19-17-19-17",
				"x-x-19-20-19-20",
			},
		},
	},
}
