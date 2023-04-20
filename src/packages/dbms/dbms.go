package dbms

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func ReadFile(FileName string) ([]byte, error) {
	File, err := os.Open(FileName)
	defer File.Close()

	FileSystem, err := os.Stat(FileName)

	if err != nil {
		return []byte{}, err
	}

	ContentsBuffer := make([]byte, FileSystem.Size())

	_, err = File.Read(ContentsBuffer)

	if err != nil {
		return []byte{}, err
	}

	return ContentsBuffer, err
}

func ReadTable(TableName string) ([]byte, error) {
	FileContents, err := ReadFile(TableName)

	// Check if the file contents can be initialised as an object
	if !json.Valid(FileContents) || err != nil {
		return []byte{}, err
	}

	return FileContents, nil
}

func AppendDataToTable(TableName string) error {
	TableContents, err := ReadFile(TableName)

	err = json.Unmarshal(TableContents, &CUSTOMSTRUCT)
	fmt.Print(CUSTOMSTRUCT)

	if err != nil {
		log.Fatal(err)
	}
	return err
}

func WriteTable(TableName string, Data any, DataType string) {
	switch DataType {
	case "UserInfo":

	}
}
