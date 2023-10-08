package shorten

type RequestBody struct {
	LongUrl   string `json:"longUrl" binding:"required,url"`
	ShortPath string `json:"shortPath" binding:"required,alphanum"` // TODO: generate short path on server side
}
