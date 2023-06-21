package dto

type Response struct {
	Status bool  `json:"status"`
	Code   int16 `json:"code"`
	Data   Data  `json:"data"`
}

type Data struct {
	Meta string `json:"meta"`
}
