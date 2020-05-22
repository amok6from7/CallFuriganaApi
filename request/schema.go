package request

type FuriganaApiRequest struct {
	AppId      string `json:"app_id"`
	Sentence   string `json:"sentence"`
	OutputType string `json:"output_type"`
}

type FuriganaApiResponse struct {
	Converted string `json:"converted"`
}
