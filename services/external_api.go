package services

import (
	"encoding/json"
	"io"
	"net/http"
)

func FetchAge(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io?name=" + name)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Age int `json:"age"`
	}
	json.NewDecoder(resp.Body).Decode(&result)
	return result.Age, nil
}

func FetchGender(name string) (string, error) {
	resp, err := http.Get("https://api.genderize.io?name=" + name)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var result struct {
		Gender string `json:"gender"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	return result.Gender, nil
}

func FetchNationality(name string) (string, error) {
	resp, err := http.Get("https://api.nationalize.io?name=" + name)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	if len(result.Country) > 0 {
		return result.Country[0].CountryID, nil
	}
	return "", nil
}
