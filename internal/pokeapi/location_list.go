package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	defer timeTrack(time.Now(), "ListLocations")
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationsRes := Locations{}
		err := json.Unmarshal(val, &locationsRes)
		if err != nil {
			return Locations{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Locations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, dat)

	locationsRes := Locations{}
	if err := json.Unmarshal(dat, &locationsRes); err != nil {
		return Locations{}, err
	}
	return locationsRes, nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
