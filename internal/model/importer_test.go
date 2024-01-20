package model

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	t.Run("Valid Data", func(t *testing.T) {
		// Test Import function with valid ChordsDbGuitarImport data
		validData := []byte(`{"keys":["C"],"chords":{"C":[{"key":"C","suffix":"maj","positions":[]}]}}`)

		tmpfile, err := os.CreateTemp("", "testChordsDbGuitarImport*.json")
		assert.NoError(t, err, "Error creating temporary file")
		defer os.Remove(tmpfile.Name())
		defer tmpfile.Close()
		tmpfile.Write(validData)

		chordsDbGuitarImporter := Import[ChordsDbGuitarImport]{
			Path: tmpfile.Name(),
		}

		importedData, err := chordsDbGuitarImporter.Import()
		assert.NoError(t, err, "Unexpected error during Import")
		assert.NotNil(t, importedData, "Imported data is nil")
	})

	t.Run("Invalid Path", func(t *testing.T) {
		// Test Import function with invalid path
		chordsDbGuitarImporter := Import[ChordsDbGuitarImport]{
			Path: "nonexistent/path.json",
		}
		_, err := chordsDbGuitarImporter.Import()
		assert.Error(t, err, "Expected error for invalid path")
	})

	t.Run("Invalid JSON", func(t *testing.T) {
		// Test Import function with invalid JSON
		invalidData := []byte(`invalid json`)

		tmpfile, err := os.CreateTemp("", "testChordsDbGuitarImportInvalid*.json")
		assert.NoError(t, err, "Error creating temporary file")
		defer os.Remove(tmpfile.Name())
		defer tmpfile.Close()
		tmpfile.Write(invalidData)

		chordsDbGuitarImporter := Import[ChordsDbGuitarImport]{
			Path: tmpfile.Name(),
		}
		_, err = chordsDbGuitarImporter.Import()
		assert.Error(t, err, "Expected error for invalid JSON")
	})

	t.Run("JSON Unmarshal Error", func(t *testing.T) {
		// Test Import function with JSON unmarshal error
		invalidJSON := []byte(`{"keys":["C"],"chords":{"C":[{"key":"C","suffix":"maj","positions":[]}]},}`)

		tmpfile, err := os.CreateTemp("", "testChordsDbGuitarImportInvalidJSON*.json")
		assert.NoError(t, err, "Error creating temporary file")
		defer os.Remove(tmpfile.Name())
		defer tmpfile.Close()
		tmpfile.Write(invalidJSON)

		chordsDbGuitarImporter := Import[ChordsDbGuitarImport]{
			Path: tmpfile.Name(),
		}
		_, err = chordsDbGuitarImporter.Import()
		assert.Error(t, err, "Expected error for invalid JSON unmarshal")
	})

	t.Run("ReadAll Error", func(t *testing.T) {
		// Test Import function with ReadAll error
		tmpfile, err := os.CreateTemp("", "testChordsDbGuitarImportReadAllError*.json")
		assert.NoError(t, err, "Error creating temporary file")
		defer os.Remove(tmpfile.Name())
		defer tmpfile.Close()

		// Make the file unreadable to simulate the ReadAll error
		file, err := os.Open(tmpfile.Name())
		assert.NoError(t, err, "Error opening temporary file")
		file.Close() // Close the file to release the handle
		os.Chmod(tmpfile.Name(), 0200)

		chordsDbGuitarImporter := Import[ChordsDbGuitarImport]{
			Path: tmpfile.Name(),
		}
		_, err = chordsDbGuitarImporter.Import()
		assert.Error(t, err, "Expected error for io.ReadAll")
	})
}
