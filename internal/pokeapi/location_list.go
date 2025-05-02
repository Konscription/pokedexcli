package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "location-area"
	if pageURL != nil {
		url = *pageURL
	}
	// check if url(key) exists in cache
	if cachedData, found := c.pokecache.Get(url); found {
		fmt.Println("Cache hit")
		locationsResp := RespShallowLocations{}
		err := json.Unmarshal(cachedData, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}
		return locationsResp, nil
	}
	// if not found, make a request to the API
	// and add the response to the cache
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	c.pokecache.Add(url, data)
	locationsResp := RespShallowLocations{}

	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResp, nil
}
