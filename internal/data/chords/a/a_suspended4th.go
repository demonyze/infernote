package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var ASuspended4th model.Chord = model.Chord{
	Id:               "asuspended4th",
	Type:             model.Suspended4th,
	AlternativeNames: data.AlternativeNames[model.Suspended4th],
	Abbreviation:     data.Abbreviations[model.Suspended4th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.D],
		notes.Notes[notes.E],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-0-2-3-0",
				"x-0-2-2-3-0",
				"5-5-7-7-5-5",
				"5-7-7-7-5-5",
				"17-17-19-19-17-17",
				"17-19-19-19-17-17",
			},
		},
	},
}
