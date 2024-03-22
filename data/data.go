package data

import (
	"bufio"
	"embed"
	"fmt"
	"io/fs"
)

//go:embed *.txt
var dataFolder embed.FS

func ReadAllData(day int) ([]byte, error) {
	path := fmt.Sprintf("day%d.txt", day)
	data, err := dataFolder.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", path, err)
	}
	return data, nil
}

func GetScanner(day int) (*bufio.Scanner, error) {
	file, err := getDayFile(day)
	if err != nil {
		return nil, err
	}
	return bufio.NewScanner(file), nil
}

func getDayFile(day int) (fs.File, error) {
	path := fmt.Sprintf("day%d.txt", day)
	file, err := dataFolder.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file %s: %w", path, err)
	}
	return file, nil
}
