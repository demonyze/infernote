package data

import "github.com/demonyze/infernote/internal/model"

var Abbreviations map[model.ChordType]model.AbbreviationType = map[model.ChordType]map[string][]string{
	model.Major: map[model.Language][]string{
		model.EN: {"major"},
	},
}
