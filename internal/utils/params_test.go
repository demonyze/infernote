package utils

import (
	"testing"

	"github.com/demonyze/infernote/internal/constants"
	"github.com/demonyze/infernote/internal/model"
	"github.com/stretchr/testify/assert"
)

func TestGetParams(t *testing.T) {
	t.Run("Success with all params", func(t *testing.T) {
		args := []string{"--lang=de", "--name=chords.json", "--path=mypath"}
		expectedParams := model.InfernoteParams{
			Language: "de",
			FileName: "chords.json",
			FilePath: "mypath",
		}
		params := GetParams(args)
		assert.Equal(t, expectedParams, params)
	})
	t.Run("Success with some params", func(t *testing.T) {
		args := []string{"--lang=de", "--namechords.json", "--path=mypath"}
		expectedParams := model.InfernoteParams{
			Language: "de",
			FileName: constants.FileName,
			FilePath: "mypath",
		}
		params := GetParams(args)
		assert.Equal(t, expectedParams, params)
	})
	t.Run("Success with no params", func(t *testing.T) {
		args := []string{}
		expectedParams := model.InfernoteParams{
			Language: constants.Language,
			FileName: constants.FileName,
			FilePath: constants.FilePath,
		}
		params := GetParams(args)
		assert.Equal(t, expectedParams, params)
	})
}

func TestExtractNamedParam(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		param := "--name=test"
		expected := model.NamedParam{
			Key:   "name",
			Value: "test",
		}
		namedParam, err := ExtractNamedParam(param)
		assert.NoError(t, err)
		assert.Equal(t, expected, namedParam)
	})

	t.Run("Error - invalid dash syntax", func(t *testing.T) {
		param := "name=test"
		expected := model.NamedParam{}
		namedParam, err := ExtractNamedParam(param)
		assert.Error(t, err)
		assert.Equal(t, expected, namedParam)
	})

	t.Run("Error - invalid assignment syntax", func(t *testing.T) {
		param := "--name:test"
		expected := model.NamedParam{}
		namedParam, err := ExtractNamedParam(param)
		assert.Error(t, err)
		assert.Equal(t, expected, namedParam)
	})
}
