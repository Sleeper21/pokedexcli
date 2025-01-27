package pokeapi

import (
	"encoding/json"
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

	//Unmarshal the data
	areas := Locations{}
	err = json.Unmarshal(data, &areas)
	if err != nil {
		return Locations{}, err
	}

	return areas, nil
}
