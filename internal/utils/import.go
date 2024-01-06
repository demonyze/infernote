package utils

import (
	"encoding/json"
	"io"
	"os"

	"github.com/demonyze/infernote/internal/model"
)

type importType = model.ChordsDbGuitarImport

func Import[I importType](path string) (I, error) {
	var dataImport I

	jsonFile, err := os.Open(path)
	if err != nil {
		return dataImport, err
	}
	defer jsonFile.Close()

	jsonBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		return dataImport, err
	}

	unmarshalError := json.Unmarshal(jsonBytes, &dataImport)
	if unmarshalError != nil {
		return dataImport, unmarshalError
	}

	return dataImport, nil
}
