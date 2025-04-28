package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func commandMapb(cfg *config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	firstURL := cfg.Previous[0]

	urlParsed, err := url.Parse(firstURL)
	if err != nil {
		return err
	}
	urlPath := strings.Split(urlParsed.Path, "/")
	idString := urlPath[len(urlPath)-2]
	id, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	cfg.Next = cfg.Previous

	for i := 0; i < 20; i++ {
		fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", id-20+i)
		cfg.Next[i] = fullURL
	}

	for i := 0; i < 20; i++ {
		fullURL := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%d/", cfg.Next)
		res, err := http.Get(fullURL)
		if err != nil {
			return err
		}
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return err
		}
		defer res.Body.Close()
		if res.StatusCode > 299 {
			return fmt.Errorf("response failed with status code: %d and\nbody: %s", res.StatusCode, data)
		}
		var locationArea localArea
		if err := json.Unmarshal(data, &locationArea); err != nil {
			return err
		}
		fmt.Println(locationArea.Name)
	}
	return nil
}
