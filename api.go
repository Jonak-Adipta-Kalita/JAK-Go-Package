package jak_go_package

import (
	"fmt"
	"net/http"
)

const BASE_URL string = "https://jak_api.p.rapidapi.com"

type Api struct {
	RapidAPIKey string
}

func (x Api) GetJAK() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", BASE_URL + "/jak", nil)

	req.Header.Set("X-RapidAPI-Host", "jak_api.p.rapidapi.com")
	req.Header.Set("X-RapidAPI-Key", x.RapidAPIKey)

	res, _ := client.Do(req)

	fmt.Print(res.Body)
}
