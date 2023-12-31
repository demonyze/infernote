package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A7thSuspended4th model.Chord = model.Chord{
	Id:               "a7thsuspended4th",
	Type:             model.SeventhSuspended4th,
	AlternativeNames: data.AlternativeNames[model.SeventhSuspended4th],
	Abbreviation:     data.Abbreviations[model.SeventhSuspended4th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.D],
		notes.Notes[notes.E],
		notes.Notes[notes.G],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-0-0-3-0",
				"x-0-2-0-3-0",
				"x-0-2-0-3-3",
				"x-0-2-2-3-3",
				"5-5-5-7-5-5",
				"5-7-5-7-5-5",
				"17-17-17-19-17-17",
				"17-19-17-19-17-17",
			},
		},
	},
}
