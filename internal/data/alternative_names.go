package data

import "github.com/demonyze/infernote/internal/model"

var AlternativeNames = map[model.AlternativeName]model.AlternativeNamesType{
	model.Chord5th: map[model.Language][]string{model.EN: {"power chord"}},
	model.Chord7th: map[model.Language][]string{model.EN: {"Dominant 7th", "Major Minor 7th"}},
}
