package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMinor6th model.Chord = model.Chord{
	Id:               "aminor6th",
	Type:             model.Minor6th,
	AlternativeNames: data.AlternativeNames[model.Minor6th],
	Abbreviation:     data.Abbreviations[model.Minor6th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.E],
		notes.Notes[notes.FSharp],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-2-1-2",
				"5-7-7-5-7-5",
				"x-x-7-9-7-8",
				"x-12-10-11-10-12",
				"17-19-19-17-19-17",
				"x-x-19-21-19-20",
			},
		},
	},
}
