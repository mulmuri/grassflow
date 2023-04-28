package io

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)




func Write(name string, content string, path string) error {
	filePath := filepath.Join(path, name)

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %v\n", err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(content)
	if err != nil {
		return fmt.Errorf("Failed to write content: %v\n", err)
	}

	err = writer.Flush()
	if err != nil {
		return fmt.Errorf("Failed to flush writer: %v\n", err)
	}

	return nil
}