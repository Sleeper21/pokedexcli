package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// List Locations
func (c *Client) GetLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	// Check cached data first
	//then we check if the URL is in the cache
	cachedData, exists := c.pokeLocalCache.Get(url)
	if exists {
		// Unmarshal the data stored in the cache
		areas := Locations{}
		err := json.Unmarshal(cachedData, &areas)
		if err != nil {
			return Locations{}, err
		}
		fmt.Println("Data retrieved from cache")
		return areas, nil
	}

	// If not in cache, make the request
	fmt.Println("Data retrieved from API")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, data)
	}
	if err != nil {
		return Locations{}, err
	}

	// Add the data to the cache
	c.pokeLocalCache.Add(url, data)

	//Unmarshal the data
	areas := Locations{}
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return Locations{}, err
	}

	return areas, nil
}
