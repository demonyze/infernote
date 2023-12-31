package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMaj9th model.Chord = model.Chord{
	Id:               "amaj9th",
	Type:             model.Major9th,
	AlternativeNames: data.AlternativeNames[model.Major9th],
	Abbreviation:     data.Abbreviations[model.Major9th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
		notes.Notes[notes.GSharp],
		notes.Notes[notes.B],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-4-6-4-5-4",
				"17-16-18-16-17-16",
			},
		},
	},
}
