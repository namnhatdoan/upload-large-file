package payloads


type CommonResponse struct  {
	Code int `json:"code"`
	Message string `json:"message"`

}

type PageFilter struct {
	Cursor int `json:"cursor"`
	Limit int `json:"limit"`
}
