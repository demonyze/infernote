package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/demonyze/infernote/internal/model"
)

func ImportChordsDbGuitar() (model.ChordsDbGuitarImport, error) {

	var chordDbGuitarImport model.ChordsDbGuitarImport

	guitarChordsFile, err := os.Open("assets/data.json")
	if err != nil {
		return chordDbGuitarImport, err
	}
	defer guitarChordsFile.Close()

	guitarChordsBytes, err := io.ReadAll(guitarChordsFile)
	if err != nil {
		return chordDbGuitarImport, err
	}

	unmarshalError := json.Unmarshal(guitarChordsBytes, &chordDbGuitarImport)
	if unmarshalError != nil {
		return chordDbGuitarImport, unmarshalError
	}

	return chordDbGuitarImport, nil
}
