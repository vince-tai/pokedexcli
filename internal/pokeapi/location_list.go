package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocations(pageURL *string) (ResLocations, error) {
	defer timeTrack(time.Now(), "ListLocations")
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationRes := ResLocations{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return ResLocations{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResLocations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return ResLocations{}, err
	}

	locationRes := ResLocations{}
	if err := json.Unmarshal(dat, &locationRes); err != nil {
		return ResLocations{}, err
	}

	return locationRes, nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
