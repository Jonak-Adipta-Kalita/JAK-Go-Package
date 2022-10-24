package jak_go_package

type JAK struct {
	Hobby        interface{}
	Games        []Game
	Pictures     []Picture
	FavFood      []FavFood     `json:"fav_food"`
	SocialMedias []SocialMedia `json:"social_medias"`
}

type Game struct {
	Id          int
	Platform    string
	Value       string
	Description string
	ImageURL    string
	Link        string
}

type Picture struct {
	Id     int
	Name   string
	Height string
	Width  string
	Image  string
}

type FavFood struct {
	Id          int
	Value       string
	ImageURL    string
	Ingredients []string
}

type SocialMedia struct {
	Id       int
	Value    string
	ImageURL string
	Link     string
	Username string
}
