package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedUrl, exists := c.pokeCache.Get(url)
	if exists {
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedUrl, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	c.pokeCache.Add(url, dat)

	locationsResp := RespShallowLocations{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}

func (c *Client) ListPokemonEncounters(pageURL string) (RespExploreLocation, error) {
	url := baseURL + "/location-area/"
	if pageURL != "" {
		url = url + pageURL
	}

	cachedUrl, exists := c.pokeCache.Get(url)
	if exists {
		pokemonResp := RespExploreLocation{}
		err := json.Unmarshal(cachedUrl, &pokemonResp)
		if err != nil {
			return RespExploreLocation{}, err
		}

		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExploreLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExploreLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExploreLocation{}, err
	}

	c.pokeCache.Add(url, dat)

	pokemonResp := RespExploreLocation{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespExploreLocation{}, err
	}

	return pokemonResp, nil
}
