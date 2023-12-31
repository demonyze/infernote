package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A6thAdd9 model.Chord = model.Chord{
	Id:               "a6thadd9",
	Type:             model.SixthAdd9,
	AlternativeNames: data.AlternativeNames[model.SixthAdd9],
	Abbreviation:     data.Abbreviations[model.SixthAdd9],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
		notes.Notes[notes.FSharp],
		notes.Notes[notes.B],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-4-4-4-5-5",
				"x-12-11-11-12-12",
				"17-16-16-16-17-17",
			},
		},
	},
}
