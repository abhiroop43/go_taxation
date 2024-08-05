package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file, it may already have been closed")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	if err != nil {
		//_ = file.Close()
		return nil, errors.New("failed to open file")
	}

	//_ = file.Close()
	return lines, nil
}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("error closing file, it may already have been closed")
		}
	}(file)

	err = json.NewEncoder(file).Encode(data)

	if err != nil {
		//_ = file.Close()
		return errors.New("failed to write data to file")
	}

	//_ = file.Close()

	return nil
}

func New(inputPath, outputPath string) FileManager {
	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
