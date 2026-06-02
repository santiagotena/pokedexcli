package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func FetchLocations(url string) (*string, *string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching locations")
		return nil, nil, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return nil, nil, err
	}

	locationsArea := LocationArea{}
	err = json.Unmarshal(dat, &locationsArea)
	if err != nil {
		fmt.Println("Error marshalling response body")
		return nil, nil, err
	}

	for _, result := range locationsArea.Results {
		fmt.Println(result.Name)
	}

	prev := locationsArea.Previous
	next := locationsArea.Next

	return prev, next, nil
}
