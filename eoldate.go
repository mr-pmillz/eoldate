package eoldate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const (
	CurrentVersion = `v0.0.8`
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

// UnmarshalJSON implements the json.Unmarshaler interface
func (p *Product) UnmarshalJSON(data []byte) error {
	type ProductAlias Product
	alias := &struct {
		*ProductAlias
		AdditionalFields map[string]interface{} `json:"-"`
	}{
		ProductAlias: (*ProductAlias)(p),
	}

	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	p.AdditionalFields = make(map[string]interface{})
	for k, v := range raw {
		switch k {
		case "cycle", "releaseDate", "eol", "latest", "link", "latestReleaseDate", "lts", "support", "extendedSupport", "minJavaVersion", "supportedPHPVersions":
			// These fields are already handled by the struct
		default:
			p.AdditionalFields[k] = v
		}
	}

	return nil
}

// GetSupportedPHPVersions returns the supported PHP versions as a string
func (p *Product) GetSupportedPHPVersions() string {
	switch v := p.SupportedPHPVersions.(type) {
	case string:
		return v
	case float64:
		return fmt.Sprintf("%.1f", v)
	default:
		return "N/A"
	}
}

// IsVersionSupported checks if the given version is supported in any of the product cycles
func (p Products) IsVersionSupported(version float64) (bool, error) {
	for _, product := range p {
		productCycle, err := strconv.ParseFloat(product.Cycle, 64)
		if err != nil {
			continue // Skip invalid cycles
		}

		if productCycle == version {
			eolDate, err := product.GetEOLDate()
			if err != nil {
				return false, err
			}

			return time.Now().Before(eolDate), nil
		}
	}
	return false, fmt.Errorf("version %.1f not found in any product cycle", version)
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
	data, err := c.Get(fmt.Sprintf("%s.json", product))
	if err != nil {
		return nil, err
	}

	var products Products
	err = json.Unmarshal(data, &products)
	return products, err
}

// GetAllProducts fetches the end-of-life information for all products.
func (c *Client) GetAllProducts() (AllProducts, error) {
	data, err := c.Get("all.json")
	if err != nil {
		return nil, err
	}

	var all AllProducts
	err = json.Unmarshal(data, &all)
	return all, err
}
