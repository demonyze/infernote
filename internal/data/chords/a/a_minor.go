package data

import (
	"github.com/demonyze/infernote/internal/data"
	notes "github.com/demonyze/infernote/internal/data/notes"
	"github.com/demonyze/infernote/internal/model"
)

var AMinor model.Chord = model.Chord{
	Id: "amin",
	// Name:         data.ChordNames[model.AMinor],
	Type:         model.Minor,
	Abbreviation: data.Abbreviations[model.Minor],
	Root:         notes.Notes[notes.A],
	Notes: []model.Note{
		notes.Notes[notes.A],
		notes.Notes[notes.C],
		notes.Notes[notes.E],
	},
	GuitarVariations: []model.GuitarVariation{
		{
			Tuning: data.Tuning[model.Standard],
			Positions: []string{
				"x-0-2-2-1-0",
				"5-7-7-5-5-5",
				"x-12-14-14-13-12",
				"17-19-19-17-17-17",
			},
		},
	},
}
