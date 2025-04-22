package external

import (
	"encoding/json"
	"net/http"
)

// общая обёртка
func fetch(url string, dst interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(dst)
}

func FetchAge(name string) (int, error) {
	var x struct {
		Age int `json:"age"`
	}
	err := fetch("https://api.agify.io?name="+name, &x)
	return x.Age, err
}

func FetchGender(name string) (string, error) {
	var x struct {
		Gender string `json:"gender"`
	}
	err := fetch("https://api.genderize.io?name="+name, &x)
	return x.Gender, err
}

func FetchNationality(name string) (string, error) {
	var x struct {
		Country []struct {
			CountryID string `json:"country_id"`
		} `json:"country"`
	}
	err := fetch("https://api.nationalize.io?name="+name, &x)
	if len(x.Country) > 0 {
		return x.Country[0].CountryID, err
	}
	return "", err
}
