package domain

type Response struct {
	Status int `json:"status"`
	Body   any
}
