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
	T []model.Chord | map[string]model.Chord,
] struct {
	Path     string
	FileName string
}

func (ex ExportMockSuccess[T]) Export(data T) error {
	return nil
}

type ExportMockError[
	T []model.Chord | map[string]model.Chord,
] struct {
	Path     string
	FileName string
}

func (ex ExportMockError[T]) Export(data T) error {
	return errors.New("mock error")
}

type ImportMockSuccess[
	T model.ChordsDbGuitarImport | model.ChordRocksGuitarImport,
] struct {
	Path string
}

func (imp ImportMockSuccess[T]) Import() (T, error) {
	var data T
	return data, nil
}

type ImportMockError[
	T model.ChordsDbGuitarImport | model.ChordRocksGuitarImport,
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

	// Write sample data to the temporary files
	chordRocksData := []byte(`{"A": {"11": {
            "internal_suffix_name": "11th",
            "notes": "A, C♯, E, G, B, D",
            "degrees": "1, 3, 5, ♭7, 9, 11",
            "abbrevation": "11",
            "alternative_names": ""
          }}}`)
	chordsDbData := []byte(`{"keys":["C"],"chords":{"C":[{"key":"C","suffix":"maj","positions":[]}]}}`)

	chordRocksFile.Write(chordRocksData)
	chordsDbFile.Write(chordsDbData)

	FileName := "output.json"
	Path := ""

	chordsArrayExporterSuccess := ExportMockSuccess[[]model.Chord]{FileName, Path}
	chordsMapExporterSuccess := ExportMockSuccess[map[string]model.Chord]{FileName, Path}
	chordsArrayExporterError := ExportMockError[[]model.Chord]{FileName, Path}
	chordsMapExporterError := ExportMockError[map[string]model.Chord]{FileName, Path}

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

	t.Run("Successful Run - Array", func(t *testing.T) {
		runner := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			ChordsArrayExporter:      chordsArrayExporterSuccess,
		}
		err := Run("array", runner)
		assert.NoError(t, err, "Expected no error for a successful run")
	})

	t.Run("Successful Run - Map", func(t *testing.T) {
		runner := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			ChordsMapExporter:        chordsMapExporterSuccess,
		}
		err := Run("map", runner)
		assert.NoError(t, err, "Expected no error for a successful run")
	})

	t.Run("Error Runs", func(t *testing.T) {
		runner1 := model.Runner{
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			ChordsDbGuitarImporter:   chordDbGuitarImporterError,
			ChordsArrayExporter:      chordsArrayExporterSuccess,
		}
		err1 := Run("array", runner1)
		assert.Error(t, err1, "Expected error")

		runner2 := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			ChordsArrayExporter:      chordsArrayExporterError,
		}
		err2 := Run("array", runner2)
		assert.Error(t, err2, "Expected error")

		runner3 := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporterError,
			ChordRocksGuitarImporter: chordRocksGuitarImporterError,
			ChordsMapExporter:        chordsMapExporterSuccess,
		}
		err3 := Run("map", runner3)
		assert.Error(t, err3, "Expected error")

		runner4 := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporterSuccess,
			ChordRocksGuitarImporter: chordRocksGuitarImporterSuccess,
			ChordsMapExporter:        chordsMapExporterError,
		}
		err4 := Run("map", runner4)
		assert.Error(t, err4, "Expected error")
	})
}
