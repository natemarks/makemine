package input

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type MakeMineInput struct {
	Name      string `json:"name"`
	LocalUser string `json:"localUser"`
	Email     string `json:"email"`
}

func FromUrl(url string) (MakeMineInput, error) {
	var data MakeMineInput
	hClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return data, err
	}

	res, err := hClient.Do(req)
	if err != nil {
		return data, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return data, err
	}

	return data, err
}

func FromFile(mFile string) (MakeMineInput, error) {
	var data MakeMineInput
	file, _ := ioutil.ReadFile(mFile)

	err := json.Unmarshal([]byte(file), &data)

	return data, err
}
