package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocationAreas(url string) (*string, *string, []byte, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching location areas")
		return nil, nil, nil, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return nil, nil, nil, err
	}

	locationAreas := LocationAreas{}
	err = json.Unmarshal(dat, &locationAreas)
	if err != nil {
		fmt.Println("Error marshalling response body")
		return nil, nil, nil, err
	}

	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}

	prev := locationAreas.Previous
	next := locationAreas.Next

	return prev, next, dat, nil
}
