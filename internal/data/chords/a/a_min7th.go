package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMin7th model.Chord = model.Chord{
	Id:               "amin7th",
	Type:             model.Minor7th,
	AlternativeNames: data.AlternativeNames[model.Minor7th],
	Abbreviation:     data.Abbreviations[model.Minor7th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.E],
		notes.Notes[notes.G],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-0-1-0",
				"x-0-2-0-1-3",
				"x-0-2-2-1-3",
				"5-7-5-5-5-5",
				"x-x-7-9-8-8",
				"x-12-10-12-10-12",
				"x-12-14-12-13-12",
				"17-19-17-17-17-17",
				"x-x-19-21-20-20",
			},
		},
	},
}
