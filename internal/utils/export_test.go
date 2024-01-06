package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestExport(t *testing.T) {
	// create temporary directory for testing
	tmpDir, err := os.MkdirTemp("", "export_test")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	testData := []model.Chord{
		{
			Id:   "00",
			Name: "Test",
		},
	}

	exportParams := ExportParams[[]model.Chord]{
		Path:     tmpDir,
		FileName: "test_export.json",
		Data:     testData,
	}

	err = Export(exportParams)
	assert.NoError(t, err, "Export should not return an error")

	filePath := filepath.Join(tmpDir, "test_export.json")
	assert.FileExists(t, filePath, "Exported file should exist")

	fileContent, err := os.ReadFile(filePath)
	assert.NoError(t, err, "Failed to read exported file")

	var exportedData []model.Chord
	err = json.Unmarshal(fileContent, &exportedData)
	assert.NoError(t, err, "Failed to unmarshal exported data")

	t.Log("Export test passed successfully")
}
