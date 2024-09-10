package eoldate

import (
	"encoding/json"
	"fmt"
	"github.com/Masterminds/semver/v3"
	"io"
	"net/http"
	"slices"
	"time"
)

const (
	CurrentVersion = `v1.0.1`
	EOLBaseURL     = "https://endoflife.date/api"
	NotAvailable   = "N/A"
)

// Options ...
type Options struct {
	Tech    string
	Output  string
	Version bool
	GetAll  bool
}

// Product represents the structure of the JSON data
type Product struct {
	Cycle                string                 `json:"cycle,omitempty"`
	ReleaseDate          string                 `json:"releaseDate,omitempty"`
	EOL                  interface{}            `json:"eol,omitempty"`
	Latest               string                 `json:"latest,omitempty"`
	Link                 string                 `json:"link,omitempty"`
	LatestReleaseDate    string                 `json:"latestReleaseDate,omitempty"`
	LTS                  interface{}            `json:"lts,omitempty"`
	Support              interface{}            `json:"support,omitempty"`
	ExtendedSupport      interface{}            `json:"extendedSupport,omitempty"`
	MinJavaVersion       *float64               `json:"minJavaVersion,omitempty"`
	SupportedPHPVersions interface{}            `json:"supportedPHPVersions,omitempty"`
	AdditionalFields     map[string]interface{} `json:"-"`
}

// IsSupportedSoftwareVersion checks if a given software version is supported
func (c *Client) IsSupportedSoftwareVersion(softwareName string, version string) (bool, error) {
	softwareReleaseData, err := c.GetProduct(softwareName)
	if err != nil {
		return false, LogError(err)
	}
	isSupported, err := softwareReleaseData.IsVersionSupported(version)
	if err != nil {
		return false, LogError(err)
	}
	return isSupported, nil
}

// IsVersionSupported checks if the given version is supported in any of the product cycles
func (p Products) IsVersionSupported(versionStr string) (bool, error) {
	version, err := semver.NewVersion(versionStr)
	if err != nil {
		return false, fmt.Errorf("invalid version string: %s", versionStr)
	}

	for _, product := range p {
		constraint, err := semver.NewConstraint(product.Cycle)
		if err != nil {
			// If the cycle is not a valid constraint, skip it
			continue
		}

		if constraint.Check(version) {
			eolDate, err := product.GetEOLDate()
			if err != nil {
				return false, err
			}
			return time.Now().Before(eolDate), nil
		}
	}
	return false, fmt.Errorf("version %s not found in any product cycle", versionStr)
}

// GetEOLDate returns the end-of-life date for the product
func (p *Product) GetEOLDate() (time.Time, error) {
	switch eol := p.EOL.(type) {
	case string:
		formats := []string{
			"2006-01-02",
			"2006-01",
			"2006",
		}
		for _, format := range formats {
			if t, err := time.Parse(format, eol); err == nil {
				return t, nil
			}
		}
		return time.Time{}, fmt.Errorf("unable to parse EOL date: %s", eol)
	case bool:
		if eol {
			return time.Now().AddDate(-1, 0, 0), nil // Assume EOL was a year ago if true
		}
		return time.Now().AddDate(100, 0, 0), nil // Assume far in the future if false
	default:
		return time.Time{}, fmt.Errorf("unexpected EOL type: %T", p.EOL)
	}
}

type AllProducts []string

// Client is the API client for the endoflife.date API.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new API client with the given base URL.
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    EOLBaseURL,
	}
}

// Get fetches data from a given endpoint.
func (c *Client) Get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.baseURL, endpoint)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

// Products represents a collection of Product
type Products []Product

// GetProduct fetches the end-of-life information for a specific product.
func (c *Client) GetProduct(product string) (Products, error) {
	allProducts, err := c.CacheTechnologies()
	if err != nil {
		return nil, err
	}
	if slices.Contains(allProducts, product) {
		var products Products
		productCache, err := readCache(product)
		if err != nil {
			return nil, err
		}
		if productCache != nil {
			err = json.Unmarshal(productCache, &products)
			return products, err
		}
		data, err := c.Get(fmt.Sprintf("%s.json", product))
		if err != nil {
			return nil, err
		}

		if err = json.Unmarshal(data, &products); err != nil {
			return nil, err
		}
		if err = writeCache(product, data); err != nil {
			return nil, err
		}
		return products, err
	} else {
		return nil, fmt.Errorf("product %s not found", product)
	}
}

// GetAllProducts fetches the end-of-life information for all products.
func (c *Client) GetAllProducts() (AllProducts, error) {
	allProductsCache, err := readAllTechnologiesCache()
	if err != nil {
		return nil, LogError(err)
	}
	if allProductsCache != nil {
		return allProductsCache, nil
	}
	data, err := c.Get("all.json")
	if err != nil {
		return nil, err
	}

	var all AllProducts
	err = json.Unmarshal(data, &all)
	return all, err
}
