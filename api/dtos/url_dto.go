package dtos

type ShortenUrlDto struct {
	Url string `json:"url"`
}

type ShortenUrlResponse struct {
	UrlPath string `json:"urlPath"`
}
