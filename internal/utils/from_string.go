package utils

import (
	"strings"

	"github.com/demonyze/infernote/internal/model"
)

func NotesfromStringNames(namesString string) []model.Note {
	namesArray := strings.Split(namesString, ",")

	var notes []model.Note
	for _, name := range namesArray {
		notes = append(notes, model.Note{Name: strings.TrimSpace(name)})
	}
	return notes
}
