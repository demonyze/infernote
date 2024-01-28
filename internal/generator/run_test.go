package generator

import (
	"errors"
	"os"
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ExportMockSuccess[
	T model.Infernote,
] struct {
	Path     string
	FileName string
}

func (ex ExportMockSuccess[T]) Export(data T) error {
	return nil
}

type ExportMockError[
	T model.Infernote,
] struct {
	Path     string
	FileName string
}

func (ex ExportMockError[T]) Export(data T) error {
	return errors.New("mock error")
}

type ImportMockSuccess[
	T model.ChordsDbGuitarImport | model.ChordRocksGuitarImport | model.LanguageImport,
] struct {
	Path string
}

func (imp ImportMockSuccess[T]) Import() (T, error) {
	var data T
	return data, nil
}

type ImportMockError[
	T model.ChordsDbGuitarImport | model.ChordRocksGuitarImport | model.LanguageImport,
] struct {
	Path string
}

func (imp ImportMockError[T]) Import() (T, error) {
	var data T
	return data, errors.New("mock error")
}

func TestRun(t *testing.T) {
	// Create temporary files for ChordRocksGuitarImport and ChordsDbGuitarImport
	chordRocksFile, err := os.CreateTemp("", "chordRocks*.json")
	require.NoError(t, err, "Error creating temporary ChordRocksGuitarImport file")
	defer os.Remove(chordRocksFile.Name())
	defer chordRocksFile.Close()

	chordsDbFile, err := os.CreateTemp("", "chordsDb*.json")
	require.NoError(t, err, "Error creating temporary ChordsDbGuitarImport file")
	defer os.Remove(chordsDbFile.Name())
	defer chordsDbFile.Close()

	languageFile, err := os.CreateTemp("", "lang*.json")
	require.NoError(t, err, "Error creating temporary language file")
	defer os.Remove(languageFile.Name())
	defer languageFile.Close()

	// Write sample data to the temporary files
	chordRocksData := []byte(`{"A": {"11": {
            "internal_suffix_name": "11th",
            "notes": "A, C♯, E, G, B, D",
            "degrees": "1, 3, 5, ♭7, 9, 11",
            "abbrevation": "11",
            "alternative_names": ""
          }}}`)
	chordsDbData := []byte(`{"keys":["C"],"chords":{"C":[{"key":"C","suffix":"maj","positions":[]}]}}`)

	langugeData := []byte(`{
		"language": "en",
		"types": {
		  "major": "Major",
		  "minor": "Minor",
		  "5": "5th",
		  "7": "7th"}}`)

	chordRocksFile.Write(chordRocksData)
	chordsDbFile.Write(chordsDbData)
	languageFile.Write(langugeData)

	FileName := "output.json"
	Path := ""

	exporterSuccess := ExportMockSuccess[model.Infernote]{FileName, Path}
	exporterError := ExportMockError[model.Infernote]{FileName, Path}

	chordRocksGuitarImporterSuccess := ImportMockSuccess[model.ChordRocksGuitarImport]{
		Path: chordRocksFile.Name(),
	}
	chordDbGuitarImporterSuccess := ImportMockSuccess[model.ChordsDbGuitarImport]{
		Path: chordsDbFile.Name(),
	}
	chordRocksGuitarImporterError := ImportMockError[model.ChordRocksGuitarImport]{
		Path: chordRocksFile.Name(),
	}
	chordDbGuitarImporterError := ImportMockError[model.ChordsDbGuitarImport]{
		Path: chordsDbFile.Name(),
	}
	languageImporterSuccess := ImportMockSuccess[model.LanguageImport]{
		Path: languageFile.Name(),
	}
	languageImporterError := ImportMockError[model.LanguageImport]{
		Path: languageFile.Name(),
	}

	t.Run("Successful Run", func(t *testing.T) {
		runner := model.Runner{
			LanguageImporter:         languageImporterSuccess,
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			InfernoteExporter:        exporterSuccess,
		}
		err := Run(runner)
		assert.NoError(t, err, "Expected no error for a successful run")
	})

	t.Run("Error Runs", func(t *testing.T) {
		runner1 := model.Runner{
			LanguageImporter:         languageImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			ChordsDbGuitarImporter:   chordDbGuitarImporterError,
		}
		err1 := Run(runner1)
		assert.Error(t, err1, "Expected error")

		runner2 := model.Runner{
			LanguageImporter:         languageImporterSuccess,
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			InfernoteExporter:        exporterError,
		}
		err2 := Run(runner2)
		assert.Error(t, err2, "Expected error")

		runner3 := model.Runner{
			LanguageImporter:         languageImporterSuccess,
			ChordsDbGuitarImporter:   chordDbGuitarImporterError,
			ChordRocksGuitarImporter: chordRocksGuitarImporterError,
			InfernoteExporter:        exporterSuccess,
		}
		err3 := Run(runner3)
		assert.Error(t, err3, "Expected error")

		runner4 := model.Runner{
			LanguageImporter:         languageImporterSuccess,
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			InfernoteExporter:        exporterError,
		}
		err4 := Run(runner4)
		assert.Error(t, err4, "Expected error")

		runner5 := model.Runner{
			LanguageImporter:         languageImporterError,
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			InfernoteExporter:        exporterSuccess,
		}
		err5 := Run(runner5)
		assert.Error(t, err5, "Expected error")
	})
}
