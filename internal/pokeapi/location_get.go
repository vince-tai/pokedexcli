package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locName string) (Location, error) {
	url := baseURL + "/location-area/" + locName

	if val, ok := c.cache.Get(url); ok {
		locationRes := Location{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Location{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	locationRes := Location{}
	if err := json.Unmarshal(dat, &locationRes); err != nil {
		return Location{}, err
	}
	return locationRes, nil
}
