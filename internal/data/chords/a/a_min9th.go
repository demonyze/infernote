package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMin9th model.Chord = model.Chord{
	Id:               "amin9th",
	Type:             model.Minor9th,
	AlternativeNames: data.AlternativeNames[model.Minor9th],
	Abbreviation:     data.Abbreviations[model.Minor9th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.E],
		notes.Notes[notes.G],
		notes.Notes[notes.B],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-7-5-5-5-7",
				"17-19-17-17-17-19",
			},
		},
	},
}
