package shorten

type RequestBody struct {
	LongURL  string `json:"longUrl"`
	ShortURL string `json:"shortUrl"`
}
