package utils

import (
	"fmt"
	"strings"

	"github.com/demonyze/infernote/internal/model"
)

func NotesfromStringNames(namesString string) []model.Note {
	var notes []model.Note = []model.Note{}
	if namesString == "" {
		return notes
	}
	namesArray := strings.Split(namesString, ",")
	for _, name := range namesArray {
		note, err := FindNoteByName(name)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		notes = append(notes, note)
	}
	return notes
}
