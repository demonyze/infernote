package fetcher

import (
	"encoding/json"
	"net/http"
)

func GetJson[T any](url string) (T, error) {
	var data T
	res, err := http.Get(url)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}
