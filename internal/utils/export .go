package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/demonyze/infernote/internal/model"
)

type exportType = []model.Chord

type ExportParams[T exportType] struct {
	Path     string
	FileName string
	Data     T
}

func Export[T exportType](exportData ExportParams[T]) error {
	chordsBytes, err := json.Marshal(exportData.Data)
	if err != nil {
		return err
	}

	createFolderError := os.MkdirAll(exportData.Path, os.ModePerm)
	if createFolderError != nil {
		return createFolderError
	}

	filePath := fmt.Sprintf("%s/%s", exportData.Path, exportData.FileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	t, err := file.Write(chordsBytes)
	if err != nil {
		return err
	}

	size := fmt.Sprintf("(%d bytes)", t)
	fmt.Println("💾 files is saved at" + " " + filePath + " " + size)

	return nil
}
