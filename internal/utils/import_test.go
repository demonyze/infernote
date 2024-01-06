package utils

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "import_test.json")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	sampleData := model.ChordsDbGuitarImport{
		Keys: []string{"A", "B", "C"},
		Chords: map[string][]model.ChordsDbGuitarChord{
			"A": {model.ChordsDbGuitarChord{
				Key:    "A",
				Suffix: "major",
				Positions: []model.ChordsDbGuitarPosition{
					{
						Frets: []int{4, 5, 4, 5, 6},
					},
				},
			}},
		},
	}

	jsonBytes, err := json.Marshal(sampleData)
	if err != nil {
		t.Fatalf("Failed to marshal sample data: %v", err)
	}

	_, err = tmpFile.Write(jsonBytes)
	if err != nil {
		t.Fatalf("Failed to write to temporary file: %v", err)
	}

	importedData, err := Import[model.ChordsDbGuitarImport](tmpFile.Name())
	assert.NoError(t, err, "Import should not return an error")

	assert.Equal(t, sampleData, importedData, "Imported data should match the sample data")

	t.Log("Import test passed successfully")
}
