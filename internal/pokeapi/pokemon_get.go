package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

//https://pokeapi.co/api/v2/pokemon/clefairy/

func (c *Client) GetPokemon(pokeName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokeName

	if val, ok := c.cache.Get(url); ok {
		pokeRes := Pokemon{}
		err := json.Unmarshal(val, &pokeRes)
		if err != nil {
			return Pokemon{}, err
		}
		return pokeRes, nil
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(url, dat)

	pokeRes := Pokemon{}
	if err := json.Unmarshal(dat, &pokeRes); err != nil {
		return Pokemon{}, err
	}
	return pokeRes, nil
}
