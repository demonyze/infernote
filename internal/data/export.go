package data

import (
	"encoding/json"
	"fmt"
	"os"
)

func ExportData(path string, fileName string) error {
	chords, err := GenerateChords()
	if err != nil {
		return err
	}

	chordsBytes, err := json.Marshal(chords)
	if err != nil {
		return err
	}

	createFolderError := os.MkdirAll(path, os.ModePerm)
	if createFolderError != nil {
		return createFolderError
	}

	filePath := fmt.Sprintf("%s/%s", path, fileName)

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
