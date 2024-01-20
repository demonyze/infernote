package generator

import (
	"os"
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type ExportMock[
	T []model.Chord | map[string]model.Chord,
] struct {
	Path     string
	FileName string
}

func (ex ExportMock[T]) Export(data T) error {
	return nil
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

	// Run the function with different scenarios
	t.Run("Successful Run - Array", func(t *testing.T) {
		chordsArrayExporter := ExportMock[[]model.Chord]{
			FileName: "output_array.json",
			Path:     "",
		}

		chordRocksGuitarImporter := model.Importer[model.ChordRocksGuitarImport]{
			Path: chordRocksFile.Name(),
		}
		chordDbGuitarImporter := model.Importer[model.ChordsDbGuitarImport]{
			Path: chordsDbFile.Name(),
		}

		runner := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporter,
			ChordRocksGuitarImporter: chordRocksGuitarImporter,
			ChordsArrayExporter:      chordsArrayExporter,
		}

		err := Run("array", runner)
		assert.NoError(t, err, "Expected no error for a successful run")
	})

	t.Run("Successful Run - Map", func(t *testing.T) {
		chordsMapExporter := ExportMock[map[string]model.Chord]{
			FileName: "output_map.json",
			Path:     "",
		}

		chordRocksGuitarImporter := model.Importer[model.ChordRocksGuitarImport]{
			Path: chordRocksFile.Name(),
		}
		chordDbGuitarImporter := model.Importer[model.ChordsDbGuitarImport]{
			Path: chordsDbFile.Name(),
		}

		runner := model.Runner{
			ChordsDbGuitarImporter:   chordDbGuitarImporter,
			ChordRocksGuitarImporter: chordRocksGuitarImporter,
			ChordsMapExporter:        chordsMapExporter,
		}

		err := Run("map", runner)
		assert.NoError(t, err, "Expected no error for a successful run")
	})

}
