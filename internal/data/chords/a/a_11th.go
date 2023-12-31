package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A11th model.Chord = model.Chord{
	Id:               "a6thadd9",
	Type:             model.SixthAdd9,
	AlternativeNames: data.AlternativeNames[model.SixthAdd9],
	Abbreviation:     data.Abbreviations[model.SixthAdd9],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
		notes.Notes[notes.G],
		notes.Notes[notes.B],
		notes.Notes[notes.D],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-5-5-6-5-7",
				"17-17-17-18-17-19",
			},
		},
	},
}
