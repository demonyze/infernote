package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A9th model.Chord = model.Chord{
	Id:               "a9th",
	Type:             model.Ninth,
	AlternativeNames: data.AlternativeNames[model.Ninth],
	Abbreviation:     data.Abbreviations[model.Ninth],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
		notes.Notes[notes.G],
		notes.Notes[notes.B],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-4-5-4-5-x",
				"5-7-5-6-5-7",
				"17-16-17-16-17-x",
				"17-19-17-18-17-19",
			},
		},
	},
}
