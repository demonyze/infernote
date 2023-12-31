package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AAugmented model.Chord = model.Chord{
	Id:               "adiminished",
	Type:             model.Augmented,
	AlternativeNames: data.AlternativeNames[model.Augmented],
	Abbreviation:     data.Abbreviations[model.Augmented],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.F],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-4-3-x-x-x",
				"17-16-15-x-x-x",
				"x-0-3-2-2-1",
				"x-x-7-6-6-5",
				"x-12-11-10-10-x",
				"x-x-x-14-14-13",
				"x-x-19-18-18-17",
			},
		},
	},
}
