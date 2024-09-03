package eoldate

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const CurrentVersion = `v0.0.3`
const EOLBaseURL = "https://endoflife.date/api"

// Options ...
type Options struct {
	Tech    string
	Output  string
	Version bool
	GetAll  bool
}

// Product represents the structure of the JSON data
type Product struct {
	Cycle             string      `json:"cycle,omitempty"`
	ReleaseDate       string      `json:"releaseDate,omitempty"`
	EOL               interface{} `json:"eol,omitempty"`
	Latest            string      `json:"latest,omitempty"`
	LatestReleaseDate string      `json:"latestReleaseDate,omitempty"`
	LTS               bool        `json:"lts,omitempty"`
	Support           string      `json:"support,omitempty"`
}

type AllProducts []string

// Client is the API client for the endoflife.date API.
type Client struct {
	httpClient *http.Client
	baseURL    string
}

// NewClient creates a new API client with the given base URL.
func NewClient(baseURL string) *Client {
	return &Client{
		httpClient: &http.Client{},
		baseURL:    baseURL,
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

// GetProduct fetches the end-of-life information for a specific product.
func (c *Client) GetProduct(product string) ([]Product, error) {
	data, err := c.Get(fmt.Sprintf("%s.json", product))
	if err != nil {
		return nil, LogError(err)
	}

	products := make([]Product, 0)
	err = json.Unmarshal(data, &products)
	return products, err
}

// GetAllProducts fetches the end-of-life information for all products.
func (c *Client) GetAllProducts() (AllProducts, error) {
	data, err := c.Get("all.json")
	if err != nil {
		return nil, LogError(err)
	}

	all := make(AllProducts, 0)
	err = json.Unmarshal(data, &all)
	return all, err
}
