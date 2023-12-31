package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMin7thFlat5th model.Chord = model.Chord{
	Id:               "amin7thFlat5th",
	Type:             model.Minor7thFlat5th,
	AlternativeNames: data.AlternativeNames[model.Minor7thFlat5th],
	Abbreviation:     data.Abbreviations[model.Minor7thFlat5th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.DSharp],
		notes.Notes[notes.G],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-6-5-5-x-x",
				"17-18-17-17-x-x",
				"x-0-1-0-1-3",
				"x-0-1-2-1-3",
				"x-x-7-8-8-8",
				"x-12-10-12-10-11",
				"x-12-13-12-13-x",
				"x-x-19-20-20-20",
			},
		},
	},
}
