package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A7th model.Chord = model.Chord{
	Id:               model.A7th,
	Name:             data.ChordNames[model.A7th],
	AlternativeNames: data.AlternativeNames[model.A7th],
	Abbreviation:     data.Abbreviations[model.Seventh],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
		notes.Notes[notes.G],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-0-2-0",
				"x-0-2-0-2-3",
				"5-7-5-6-5-5",
				"x-x-7-9-8-9",
				"x-12-14-12-14-12",
				"17-19-17-18-17-17",
				"x-x-19-21-20-21",
			},
		},
	},
}
