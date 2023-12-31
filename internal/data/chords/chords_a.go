package data

// import (
// 	"github.com/demonyze/infernote/internal/data"
// 	"github.com/demonyze/infernote/internal/model"
// )

// var AMajor model.Chord = model.Chord{
// 	Id:   1,
// 	Name: "A Major",
// 	Root: Notes[A],
// 	Notes: []model.Note{
// 		Notes[A],
// 		Notes[CSharp],
// 		Notes[E],
// 	},
// 	GuitarVariations: []model.GuitarVariation{
// 		{
// 			Tuning:    data.Tuning[model.Standard],
// 			Positions: []string{"x-0-2-2-2-0", "5-7-7-6-5-5", "x-12-14-14-14-12", "17-19-19-18-17-17"},
// 		},
// 	},
// }

// var AMinor model.Chord = model.Chord{
// 	Id:           2,
// 	Name:         "A Minor",
// 	Abbreviation: "minor",
// 	Root:         Notes[A],
// 	Notes: []model.Note{
// 		Notes[A],
// 		Notes[C],
// 		Notes[E],
// 	},
// 	GuitarVariations: []model.GuitarVariation{
// 		{
// 			Tuning:    data.Tuning[model.Standard],
// 			Positions: []string{"x-0-2-2-1-0", "5-7-7-5-5-5", "x-12-14-14-13-12", "17-19-19-17-17-17"},
// 		},
// 	},
// }

// var A5th model.Chord = model.Chord{
// 	Id:               3,
// 	Name:             "A 5th",
// 	Abbreviation:     "5",
// 	AlternativeNames: AlternativeNames[model.Chord5th],
// 	Root:             Notes[A],
// 	Notes: []model.Note{
// 		Notes[A],
// 		Notes[E],
// 	},
// 	GuitarVariations: []model.GuitarVariation{
// 		{
// 			Tuning:    data.Tuning[model.Standard],
// 			Positions: []string{"x-0-2-2-x-x", "5-7-7-x-x-x", "x-12-14-14-x-x", "17-19-19-x-x-x"},
// 		},
// 	},
// }

// var A7th model.Chord = model.Chord{
// 	Id:               4,
// 	Name:             "A 7th",
// 	Abbreviation:     "7",
// 	AlternativeNames: AlternativeNames[model.Chord7th],
// 	Root:             Notes[A],
// 	Notes: []model.Note{
// 		Notes[A],
// 		Notes[CSharp],
// 		Notes[E],
// 		Notes[G],
// 	},
// 	GuitarVariations: []model.GuitarVariation{
// 		{
// 			Tuning: data.Tuning[model.Standard],
// 			Positions: []string{
// 				"x-0-2-0-2-0", "x-0-2-0-2-3",
// 				"5-7-5-6-5-5", "x-x-7-9-8-9", "x-12-14-12-14-12",
// 				"17-19-17-18-17-17", "x-x-19-21-20-21",
// 			},
// 		},
// 	},
// }

// var AMajor7th model.Chord = model.Chord{
// 	Id:           5,
// 	Name:         "A Major 7th",
// 	Abbreviation: "maj7",
// 	Root:         Notes[A],
// 	Notes: []model.Note{
// 		Notes[A],
// 		Notes[CSharp],
// 		Notes[E],
// 		Notes[GSharp],
// 	},
// 	GuitarVariations: []model.GuitarVariation{
// 		{
// 			Tuning: data.Tuning[model.Standard],
// 			Positions: []string{"x-0-2-1-2-0", "5-7-6-6-5-5", "x-x-7-9-9-9",
// 				"x-12-14-13-14-12", "17-19-18-18-17-17", "x-x-19-21-21-21"},
// 		},
// 	},
// }
