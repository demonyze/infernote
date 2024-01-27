package model

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "export_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testData := Infernote{
		Language: EN,
		Types: map[string]ChordType{
			"maj": {
				Id:   "maj",
				Name: "Major",
			},
		},
		Chords: map[string]Chord{
			"cmaj": {
				Id:   "cmaj",
				Name: "C Major",
			},
		},
	}

	exporter := Export[Infernote]{
		FileName: "test_export.json",
		Path:     tmpDir,
	}

	err = exporter.Export(testData)
	assert.NoError(t, err, "Export should not return an error")

	filePath := filepath.Join(tmpDir, "test_export.json")
	assert.FileExists(t, filePath, "Exported file should exist")

	fileContent, err := os.ReadFile(filePath)
	assert.NoError(t, err, "Failed to read exported file")

	var exportedData Infernote
	err = json.Unmarshal(fileContent, &exportedData)
	assert.NoError(t, err, "Failed to unmarshal exported data")

	t.Log("Export test passed successfully")
}
