package jak_go_package

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const BASE_URL string = "https://jak_api.p.rapidapi.com"

func makeRequest(rapidApiKey, endpoint string) string {
	req, _ := http.NewRequest("GET", BASE_URL + endpoint, nil)

	req.Header.Add("X-RapidAPI-Key", rapidApiKey)
	req.Header.Add("X-RapidAPI-Host", "jak_api.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)
	
	if resErr != nil {
		fmt.Print(resErr.Error())
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		fmt.Print(bodyErr.Error())
	}

	return string(body)
}

type Api struct {
	RapidAPIKey string
}

func (x Api) GetJAK() string {
	return makeRequest(x.RapidAPIKey, "/jak")
}
