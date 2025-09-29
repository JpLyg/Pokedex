package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
)

const defaultLocationAreaPath = "/api/v2/location-area"

func (c *Client) doGET(url string) ([]byte, error) {
	// cache
	if b, ok := c.cache.Get(url); ok {
		fmt.Println("(cache) " + url)
		return b, nil
	}
	fmt.Println("(net) " + url)
	resp, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("GET %s: status %d", url, resp.StatusCode)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, b)
	return b, nil
}

// GetLocationAreasFirstPage fetches the first page (20 items).
func (c *Client) GetLocationAreasFirstPage() (LocationAreaList, error) {
	url := c.baseURL + defaultLocationAreaPath
	return c.getLocationAreas(url)
}

// GetLocationAreasByURL follows next/previous URLs.
func (c *Client) GetLocationAreasByURL(url string) (LocationAreaList, error) {
	return c.getLocationAreas(url)
}

func (c *Client) getLocationAreas(url string) (LocationAreaList, error) {
	var out LocationAreaList
	b, err := c.doGET(url)
	if err != nil {
		return out, err
	}
	if err := json.Unmarshal(b, &out); err != nil {
		return out, err
	}
	return out, nil
}
