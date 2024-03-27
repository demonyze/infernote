package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/demonyze/infernote/internal/data"
	"github.com/demonyze/infernote/internal/model"
)

func FindNoteByName(name string) (model.Note, error) {
	var lowerCaseName = strings.ToLower(name)
	var trimmedName = strings.TrimSpace(lowerCaseName)
	var noteName = strings.Replace(trimmedName, "♯", "#", -1)

	for _, note := range data.Notes {
		if strings.ToLower(note.Name) == noteName ||
			strings.ToLower(note.AlternativeName) == noteName {
			return note, nil
		}
	}
	message := fmt.Sprintf("note not found for name: %s %s", name, noteName)
	return model.Note{}, errors.New(message)
}
