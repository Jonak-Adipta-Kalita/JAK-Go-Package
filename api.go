package jak_go_package

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BASE_URL string = "https://jak_api.p.rapidapi.com"

type Api struct {
	RapidAPIKey string
}

func (x Api) GetBen10() (Ben10, error) {
	req, _ := http.NewRequest("GET", BASE_URL + "/ben10", nil)

	req.Header.Add("X-RapidAPI-Key", x.RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "jak_api.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)
	
	if resErr != nil {
		return Ben10{}, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return Ben10{}, bodyErr
	}

	var response Ben10
	json.Unmarshal(body, &response)

	return response, nil
}
