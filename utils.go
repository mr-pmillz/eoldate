package eoldate

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

// Exists returns whether the given file or directory exists
func Exists(path string) (bool, error) {
	if path == "" {
		return false, nil
	}
	absPath, err := ResolveAbsPath(path)
	if err != nil {
		return false, err
	}
	info, err := os.Stat(absPath)
	if err == nil {
		switch {
		case info.IsDir():
			return true, nil
		case info.Size() >= 0:
			// file exists but it's empty
			return true, nil
		}
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// WriteLines writes the lines to the given file.
func WriteLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return LogError(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		if len(line) > 0 {
			_, _ = fmt.Fprintln(w, line)
		}
	}
	return w.Flush()
}

// ReadLines reads a whole file into memory
// and returns a slice of its lines.
func ReadLines(path string) ([]string, error) {
	var lines []string
	absPath, err := ResolveAbsPath(path)
	if err != nil {
		return nil, LogError(err)
	}
	exists, err := Exists(absPath)
	if err != nil {
		return nil, LogError(err)
	}
	if !exists {
		fmt.Printf("File does not exist, cannot read lines for non-existent file: %s", absPath)
		return lines, nil
	}

	file, err := os.Open(absPath)
	if err != nil {
		return nil, LogError(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

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
