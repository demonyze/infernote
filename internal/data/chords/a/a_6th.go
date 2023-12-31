package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A6th model.Chord = model.Chord{
	Id:               "a6th",
	Type:             model.Sixth,
	AlternativeNames: data.AlternativeNames[model.Sixth],
	Abbreviation:     data.Abbreviations[model.Sixth],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
		notes.Notes[notes.FSharp],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-2-2-2",
				"5-4-4-6-5-x",
				"x-x-7-9-7-9",
				"17-16-16-18-17-x",
				"x-x-19-21-19-21",
			},
		},
	},
}
