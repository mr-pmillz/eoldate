package eoldate

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// readCache reads the cached data for a product
func readCache(product string) ([]byte, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, LogError(err)
	}
	timestamp := time.Now().Format("01-02-2006")
	cacheDir := filepath.Join(homeDir, ".config", "eoldate", "cache")
	cacheFile := fmt.Sprintf("%s/%s-%s.json", cacheDir, product, timestamp)
	if exists, err := Exists(cacheFile); err == nil && exists {
		return os.ReadFile(cacheFile)
	} else {
		return nil, nil
	}
}

// readAllTechnologiesCache ...
func readAllTechnologiesCache() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, LogError(err)
	}
	timestamp := time.Now().Format("01-02-2006")
	cacheDir := filepath.Join(homeDir, ".config", "eoldate", "cache")
	cacheAllTechFile := fmt.Sprintf("%s/all-technologies-%s.json", cacheDir, timestamp)
	if exists, err := Exists(cacheAllTechFile); err == nil && exists {
		return ReadLines(cacheAllTechFile)
	} else {
		return nil, nil
	}
}

// writeCache writes data to the cache for a product
func writeCache(product string, data []byte) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return LogError(err)
	}
	timestamp := time.Now().Format("01-02-2006")
	cacheDir := filepath.Join(homeDir, ".config", "eoldate", "cache")
	cacheFile := fmt.Sprintf("%s/%s-%s.json", cacheDir, product, timestamp)
	return os.WriteFile(cacheFile, data, 0600)
}

// CacheTechnologies caches all available technologies to choose from to a local file cache
func (c *Client) CacheTechnologies() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, LogError(err)
	}
	timestamp := time.Now().Format("01-02-2006")
	cacheDir := filepath.Join(homeDir, ".config", "eoldate", "cache")
	allTechnologiesFileCache := fmt.Sprintf("%s/all-technologies-%s.json", cacheDir, timestamp)
	if exists, err := Exists(cacheDir); err == nil && !exists {
		if err = os.MkdirAll(cacheDir, 0755); err != nil {
			return nil, LogError(err)
		}
	}

	if cacheExists, err := Exists(allTechnologiesFileCache); err == nil && cacheExists {
		return ReadLines(allTechnologiesFileCache)
	}
	allProducts, err := c.GetAllProducts()
	if err != nil {
		return nil, LogError(err)
	}
	if err = WriteLines(allProducts, allTechnologiesFileCache); err != nil {
		return nil, LogError(err)
	}

	return allProducts, nil
}
