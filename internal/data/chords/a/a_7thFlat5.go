package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var A7thFlat5 model.Chord = model.Chord{
	Id:               "a7thflat5",
	Type:             model.SeventhFlat5,
	AlternativeNames: data.AlternativeNames[model.SeventhFlat5],
	Abbreviation:     data.Abbreviations[model.SeventhFlat5],
	Root:             notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.DSharp],
		notes.Notes[notes.G],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"5-6-5-6-x-x",
				"17-18-17-18-x-x",
				"x-0-1-0-2-3",
				"x-0-1-2-2-3",
				"5-4-5-6-4-x",
				"x-x-7-8-8-9",
				"x-12-13-12-14-x",
				"17-16-17-18-16-x",
				"x-x-19-20-20-21",
			},
		},
	},
}
