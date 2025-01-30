package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetAreaEncounters(areaName string) (Encounters, error) {
	url := baseURL + "/location-area/" + areaName

	// Check cached data first
	// We check if the area name is in the Cache
	cachedArea, exists := c.pokeLocalCache.Get(areaName)
	if exists {
		// Unmarshal the data stored in cache
		areaDetails := Encounters{}
		err := json.Unmarshal(cachedArea, &areaDetails)
		if err != nil {
			return Encounters{}, err
		}
		fmt.Println("Data retrieved from cache")
		return areaDetails, nil
	}

	// If not in cache, make the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Encounters{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Encounters{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return Encounters{}, errors.New("response failed with status code: " + fmt.Sprint(resp.StatusCode) + ": " + string(data))
	}
	if err != nil {
		return Encounters{}, err
	}

	// Add the data to the cache
	c.pokeLocalCache.Add(areaName, data)

	//Unmarshal the data
	encounterDetails := Encounters{}
	err = json.Unmarshal(data, &encounterDetails)
	if err != nil {
		return Encounters{}, err
	}
	fmt.Println("Data retrieved from API")

	return encounterDetails, nil
}
