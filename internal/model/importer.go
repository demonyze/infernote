package model

import (
	"encoding/json"
	"io"
	"os"
)

type Importer[
	T ChordsDbGuitarImport | ChordRocksGuitarImport | LanguageImport,
] interface {
	Import() (T, error)
}

type Import[
	T ChordsDbGuitarImport | ChordRocksGuitarImport | LanguageImport,
] struct {
	Path string
	Data []byte
}

func (imp Import[T]) Import() (T, error) {
	var dataImport T
	var byteData = imp.Data

	if imp.Data == nil {
		jsonFile, err := os.Open(imp.Path)
		if err != nil {
			return dataImport, err
		}
		defer jsonFile.Close()
		jsonBytes, err := io.ReadAll(jsonFile)
		if err != nil {
			return dataImport, err
		}
		byteData = jsonBytes
	}

	unmarshalError := json.Unmarshal(byteData, &dataImport)
	if unmarshalError != nil {
		return dataImport, unmarshalError
	}

	return dataImport, nil
}
