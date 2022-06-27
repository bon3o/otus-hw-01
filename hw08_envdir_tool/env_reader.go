package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Environment map[string]EnvValue

var ErrInvalidFileName = errors.New("invalid file name")

type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("ReadDir: couldn't parse dir: %w", err)
	}
	environment := make(Environment)
	for _, file := range files {
		fileName := file.Name()
		if strings.Contains(fileName, "=") {
			return nil, ErrInvalidFileName
		}
		filePath := filepath.Join(dir, fileName)
		value, err := processFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("ReadDir: couldn't procese file %v: %w", fileName, err)
		}
		environment[fileName] = value
	}
	return environment, nil
}

func processFile(filePath string) (EnvValue, error) {
	value := EnvValue{}

	fileHandler, err := os.Open(filePath)
	if err != nil {
		return value, err
	}
	defer fileHandler.Close()

	fileInfo, err := fileHandler.Stat()
	if err != nil {
		return value, err
	}

	if fileInfo.Size() == 0 {
		value.NeedRemove = true
		return value, nil
	}

	text, err := readEnvValue(fileHandler)
	if err != nil {
		return value, err
	}

	return EnvValue{Value: text}, nil
}

func readEnvValue(handler *os.File) (string, error) {
	reader := bufio.NewReader(handler)
	bytesString, err := reader.ReadBytes('\n')
	if err != nil && err != io.EOF {
		return "", err
	}

	bytesString = bytes.ReplaceAll(bytesString, []byte{0x00}, []byte{0x0A})
	text := string(bytesString)
	text = strings.TrimRight(text, " \n\t")
	return text, nil
}
