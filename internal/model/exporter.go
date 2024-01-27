package model

import (
	"encoding/json"
	"fmt"
	"os"
)

type Exporter[T Infernote] interface {
	Export(data T) error
}

type Export[
	T Infernote,
] struct {
	Path     string
	FileName string
}

func (exp Export[T]) Export(data T) error {
	chordsBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}

	createFolderError := os.MkdirAll(exp.Path, os.ModePerm)
	if createFolderError != nil {
		return createFolderError
	}

	filePath := fmt.Sprintf("%s/%s", exp.Path, exp.FileName)

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
	fmt.Println("💾 File is saved at" + " " + filePath + " " + size)

	return nil
}
