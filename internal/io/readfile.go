package io

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/mulmuri/grassflow/internal/model"
)

func ReadFile(fileDir string) ([]model.File, error) {
	root := fileDir

	fileInfo := make([]model.File, 0)

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file := model.File{
			Name: info.Name(),
			Path: path,
			ModTime: info.ModTime(),
			Parent: getParent(path),
		}

		fileInfo = append(fileInfo, file)

		return nil

	}); err != nil {
		return nil, err
	}

	return fileInfo, nil
}


func getParent(path string) string {
	array := strings.Split(path, "/")
	return array[len(array)-2]
}