package model

type MaxRequest struct {
    Numbers []int `json:"numbers"`
}

type MaxResponse struct {
    Max int `json:"max"`
}
