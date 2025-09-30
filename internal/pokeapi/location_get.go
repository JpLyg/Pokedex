// go
package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	url := c.baseURL + "/api/v2/location-area/" + locationName

	if val, ok := c.cache.Get(url); ok {
		var out Location
		if err := json.Unmarshal(val, &out); err != nil {
			return Location{}, err
		}
		return out, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Location{}, err
	}

	var out Location
	if err := json.Unmarshal(dat, &out); err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)
	return out, nil
}
