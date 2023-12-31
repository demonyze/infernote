package data

import "github.com/demonyze/infernote/internal/model"

var Abbreviations map[model.ChordName]model.AbbreviationType = map[model.Abbreviation]map[string][]string{
	model.Major: map[model.Language][]string{
		model.EN: {"major"},
	},
}
