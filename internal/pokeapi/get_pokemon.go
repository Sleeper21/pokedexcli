package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name

	// Check if is in Cache
	val, exists := c.pokeLocalCache.Get(url)
	if exists {
		fmt.Println("Retrieving from Cache...")
		cachedPokemon := Pokemon{}
		err := json.Unmarshal(val, &cachedPokemon)
		if err != nil {
			return Pokemon{}, err
		}
		return cachedPokemon, nil
	}
	// If not
	// make new Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, errors.New("error: failed to get response from api")
	}
	defer resp.Body.Close()

	rawData, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return Pokemon{}, errors.New("response failed with status code: " + fmt.Sprint(resp.StatusCode) + ": " + string(rawData))
	}
	if err != nil {
		return Pokemon{}, errors.New("error: failed to read response from api")
	}

	// Add to cache
	c.pokeLocalCache.Add(url, rawData)

	// Unmarshal data to a Pokemon struct
	pokemon := Pokemon{}
	err = json.Unmarshal(rawData, &pokemon)
	if err != nil {
		return Pokemon{}, errors.New("error: failed to unmarshal data")
	}
	fmt.Println("Retrieving data from API...")

	return pokemon, nil
}
