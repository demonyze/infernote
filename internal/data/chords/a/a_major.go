package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMajor model.Chord = model.Chord{
	Id:           model.AMajor,
	Name:         data.ChordNames[model.AMajor],
	Abbreviation: data.Abbreviations[model.Major],
	Root:         notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.CSharp],
		notes.Notes[notes.E],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-2-2-0",
				"5-7-7-6-5-5",
				"x-12-14-14-14-12",
				"17-19-19-18-17-17",
			},
		},
	},
}
