package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMinMaj7th model.Chord = model.Chord{
	Id:               "aminmaj7th",
	Type:             model.MinorMajor7th,
	AlternativeNames: data.AlternativeNames[model.MinorMajor7th],
	Abbreviation:     data.Abbreviations[model.MinorMajor7th],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.E],
		notes.Notes[notes.GSharp],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-1-1-0",
				"5-7-6-5-5-5",
				"x-x-7-9-9-8",
				"x-12-14-13-13-12",
				"17-19-18-17-17-17",
				"x-x-19-21-21-20",
			},
		},
	},
}
