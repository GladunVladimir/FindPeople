package service

import (
	"encoding/json"
	"log"
	"net/http"
)

func FetchAge(name string) (int, error) {
	resp, err := http.Get("https://api.agify.io?name=" + name)
	if err != nil {
		return 0, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

	var result struct {
		Age int `json:"age"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return 0, err
	}
	return result.Age, nil
}

func FetchGender(name string) (string, error) {
	resp, err := http.Get("https://api.genderize.io?name=" + name)
	if err != nil {
		return "", err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

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

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("error closing response body: %v", err)
		}
	}()

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
