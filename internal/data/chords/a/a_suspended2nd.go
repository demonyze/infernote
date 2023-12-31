package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var ASuspended2nd model.Chord = model.Chord{
	Id:               "asuspended2nd",
	Type:             model.Suspended2nd,
	AlternativeNames: data.AlternativeNames[model.Suspended2nd],
	Abbreviation:     data.Abbreviations[model.Suspended2nd],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.B],
		notes.Notes[notes.E],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-2-0-0",
				"5-7-7-x-5-7",
				"x-12-14-14-12-12",
			},
		},
	},
}
