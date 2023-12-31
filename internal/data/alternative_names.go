package data

import "github.com/demonyze/infernote/internal/model"

var AlternativeNames = map[model.ChordType]map[model.Language][]string{
	model.Fith:    {model.EN: {"power chord"}},
	model.Seventh: {model.EN: {"Dominant 7th", "Major Minor 7th"}},
}
