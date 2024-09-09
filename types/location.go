package types

type Location struct {
	Name string `json:"name"`
	Url string `json:"url"`
}

type LocationResponse struct {
	Count int `json:"count"`
	Next string `json:"next"`
	Previous string `json:"previous"`
	Results []Location `json:"results"`
}