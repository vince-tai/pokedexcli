package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (c *Client) ListLocations(pageURL *string) (ResShallowLocations, error) {
	defer timeTrack(time.Now(), "ListLocations")
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationResp := ResShallowLocations{}
		err := json.Unmarshal(val, &locationResp)
		if err != nil {
			return ResShallowLocations{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return ResShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return ResShallowLocations{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return ResShallowLocations{}, err
	}

	locationRes := ResShallowLocations{}
	if err := json.Unmarshal(dat, &locationRes); err != nil {
		return ResShallowLocations{}, err
	}

	return locationRes, nil
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	fmt.Printf("%s took %s\n", name, elapsed)
}
