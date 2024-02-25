package schemas

type NumbersSetRequestSchema struct {
	Numbers []int `json:"numbers"`
}

type NumbersSetResponseSchema struct {
	Results map[string]int `json:"results"`
}
