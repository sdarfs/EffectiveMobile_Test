package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type respAge struct {
	Count int    `json:"count"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type respGender struct {
	Count       int     `json:"count"`
	Name        string  `json:"age"`
	Gender      string  `json:"gender"`
	Probability float32 `json:"probability"`
}

type respNation struct {
	Count   int            `json:"count"`
	Name    string         `json:"name"`
	Country []countrySlice `json:"country"`
}

type countrySlice struct {
	CountryId   string  `json:"country_id"`
	Probability float32 `json:"probability"`
}

func Age(name string) (int, error) {
	url := fmt.Sprintf("https://api.agify.io/?name=%s", name)
	var response respAge

	req, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	return response.Age, nil
}

func Gender(name string) (string, error) {
	url := fmt.Sprintf("https://api.genderize.io/?name=%s", name)
	var response respGender

	req, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Gender, nil
}

func Nationality(name string) (string, error) {
	url := fmt.Sprintf("https://api.nationalize.io/?name=%s", name)
	var response respNation

	req, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", err
	}

	return response.Country[0].CountryId, nil
}
