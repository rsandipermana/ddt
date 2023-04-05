package model

type CountRequest struct {
    Target string `json:"target"`
    Text string `json:"text"`
}

type CountResponse struct {
    Count int `json:"count"`
}
