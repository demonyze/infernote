package model

import (
	"encoding/json"
	"io"
	"os"
)

type Importer[
	T ChordsDbGuitarImport | ChordRocksGuitarImport,
] struct {
	Path string
}

func (imp Importer[T]) Import() (T, error) {
	var dataImport T

	jsonFile, err := os.Open(imp.Path)
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
