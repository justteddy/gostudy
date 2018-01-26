package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SearchComics search comics from xkcd database by id
func SearchComics(id string) (*Comics, error) {

	resp, err := http.Get(fmt.Sprintf(XkcdURL, id))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var result Comics
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
