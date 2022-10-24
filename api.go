package jak_go_package

import (
	"fmt"
	"net/http"
)

const BASE_URL string = "https://jak_api.p.rapidapi.com"

type api struct {
	rapidAPIKey string
}

func (x api) getJAK() {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", BASE_URL + "/jak", nil)

	req.Header.Set("X-RapidAPI-Host", "jak_api.p.rapidapi.com")
	req.Header.Set("X-RapidAPI-Key", x.rapidAPIKey)

	res, _ := client.Do(req)

	fmt.Print(res.Body)
}
