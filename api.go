package jak_go_package

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const BASE_URL string = "https://jak_api.p.rapidapi.com"

type JAK struct {
	Hobby []struct {
		ID    int    `json:"id"`
		Value string `json:"value"`
		Games []struct {
			ID          int    `json:"id"`
			Platform    string `json:"platform"`
			Value       string `json:"value"`
			Description string `json:"description"`
			Link        string `json:"link"`
			ImageURL    string `json:"imageURL"`
		} `json:"games,omitempty"`
		Languages []struct {
			ID       int    `json:"id"`
			Value    string `json:"value"`
			Website  string `json:"website"`
			Level    int    `json:"level"`
			ImageURL string `json:"imageURL"`
		} `json:"languages,omitempty"`
		Frameworks []struct {
			ID        int      `json:"id"`
			Value     string   `json:"value"`
			Website   string   `json:"website"`
			Languages []string `json:"languages"`
			Level     int      `json:"level"`
			ImageURL  string   `json:"imageURL"`
		} `json:"frameworks,omitempty"`
		Instruments []struct {
			ID       int    `json:"id"`
			Value    string `json:"value"`
			ImageURL string `json:"imageURL"`
		} `json:"instruments,omitempty"`
	} `json:"hobby"`
	Pictures []struct {
		ID     int    `json:"id"`
		Name   string `json:"name"`
		Height string `json:"height"`
		Width  string `json:"width"`
		Image  string `json:"image"`
	} `json:"pictures"`
	Games   []interface{} `json:"games"`
	FavFood []struct {
		ID          int      `json:"id"`
		Value       string   `json:"value"`
		ImageURL    string   `json:"imageURL"`
		Ingredients []string `json:"ingredients"`
	} `json:"fav_food"`
	SocialMedias []struct {
		ID       int    `json:"id"`
		Value    string `json:"value"`
		ImageURL string `json:"imageURL"`
		Link     string `json:"link"`
		Username string `json:"username"`
	} `json:"social_medias"`
}

type BrawlStars struct {
	Brawlers []struct {
		ID        int      `json:"id"`
		Name      string   `json:"name"`
		Image     string   `json:"image"`
		Gadget    []string `json:"gadget"`
		Starpower []string `json:"starpower"`
		Category  string   `json:"category"`
		Pins      []struct {
			Image string `json:"image"`
		} `json:"pins"`
		Sprays []struct {
			Image string `json:"image"`
		} `json:"sprays"`
	} `json:"brawlers"`
	Players struct {
		Pins []struct {
			Image string `json:"image"`
		} `json:"pins"`
		Sprays []struct {
			Image string `json:"image"`
		} `json:"sprays"`
	} `json:"players"`
}

type Ben10 struct {
	Omnitrix []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Image string `json:"image"`
	} `json:"omnitrix"`
}

type GenshinImpact struct {
	Character []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Image         string `json:"image"`
		Element       string `json:"element"`
		Weapon        string `json:"weapon"`
		SpecialDish   string `json:"special_dish"`
		Sex           string `json:"sex"`
		Birthday      string `json:"birthday"`
		Constellation string `json:"constellation"`
		Nation        string `json:"nation"`
	} `json:"character"`
	Elements []struct {
		ID            int    `json:"id"`
		Name          string `json:"name"`
		Image         string `json:"image"`
		Archon        string `json:"archon"`
		StatusApplies string `json:"status_applies"`
	} `json:"elements"`
}

type MughalEmpire struct {
	Description string `json:"description"`
	Kings       []struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		GivenName   string `json:"givenName"`
		Image       string `json:"image"`
		Description string `json:"description"`
		Reigned     string `json:"reigned"`
		Wives       []struct {
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Image       string `json:"image"`
			Born        string `json:"born"`
			Died        string `json:"died"`
			Description string `json:"description"`
		} `json:"wives"`
	} `json:"kings"`
}

type Api struct {
	RapidAPIKey string
}

func (x Api) GetJAK() (JAK, error) {
	req, _ := http.NewRequest("GET", BASE_URL + "/jak", nil)

	req.Header.Add("X-RapidAPI-Key", x.RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "jak_api.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)
	
	if resErr != nil {
		return JAK{}, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return JAK{}, bodyErr
	}

	var response JAK
	json.Unmarshal(body, &response)

	return response, nil
}

func (x Api) GetBrawlStars() (BrawlStars, error) {
	req, _ := http.NewRequest("GET", BASE_URL + "/brawlStars", nil)

	req.Header.Add("X-RapidAPI-Key", x.RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "jak_api.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)
	
	if resErr != nil {
		return BrawlStars{}, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return BrawlStars{}, bodyErr
	}

	var response BrawlStars
	json.Unmarshal(body, &response)

	return response, nil
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

func (x Api) GetGenshinImpact() (GenshinImpact, error) {
	req, _ := http.NewRequest("GET", BASE_URL + "/genshinImpact", nil)

	req.Header.Add("X-RapidAPI-Key", x.RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "jak_api.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)
	
	if resErr != nil {
		return GenshinImpact{}, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return GenshinImpact{}, bodyErr
	}

	var response GenshinImpact
	json.Unmarshal(body, &response)

	return response, nil
}

func (x Api) GetMughalEmpire() (MughalEmpire, error) {
	req, _ := http.NewRequest("GET", BASE_URL + "/mughalEmpire", nil)

	req.Header.Add("X-RapidAPI-Key", x.RapidAPIKey)
	req.Header.Add("X-RapidAPI-Host", "jak_api.p.rapidapi.com")

	res, resErr := http.DefaultClient.Do(req)
	
	if resErr != nil {
		return MughalEmpire{}, resErr
	}

	defer res.Body.Close()
	body, bodyErr := ioutil.ReadAll(res.Body)

	if bodyErr != nil {
		return MughalEmpire{}, bodyErr
	}

	var response MughalEmpire
	json.Unmarshal(body, &response)

	return response, nil
}
