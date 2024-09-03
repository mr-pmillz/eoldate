package eoldate

import (
	"encoding/json"
	"github.com/gocarina/gocsv"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// WriteStructToJSONFile ...
func WriteStructToJSONFile(data interface{}, outputFile string) error {
	outputFileDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputFileDir, 0750); err != nil {
		return LogError(err)
	}

	f, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return LogError(err)
	}

	if err = os.WriteFile(outputFile, f, 0644); err != nil { //nolint:gosec
		return LogError(err)
	}
	return nil
}

// WriteStructToCSVFile ...
func WriteStructToCSVFile(data interface{}, outputFile string) error {
	outputFileDir := filepath.Dir(outputFile)
	if err := os.MkdirAll(outputFileDir, 0750); err != nil {
		return LogError(err)
	}

	file, err := os.OpenFile(outputFile, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return LogError(err)
	}
	defer file.Close()

	err = gocsv.MarshalFile(data, file)
	if err != nil {
		return LogError(err)
	}

	return nil
}

// WriteStringToFile writes a string to a file
func WriteStringToFile(outputFile, data string) error {
	out, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer out.Close()
	if _, err = out.WriteString(data); err != nil {
		return LogError(err)
	}

	return nil
}

// ResolveAbsPath ...
func ResolveAbsPath(path string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return path, LogError(err)
	}

	dir := usr.HomeDir
	if path == "~" {
		path = dir
	} else if strings.HasPrefix(path, "~/") {
		path = filepath.Join(dir, path[2:])
	}

	path, err = filepath.Abs(path)
	if err != nil {
		return path, LogError(err)
	}

	return path, nil
}
