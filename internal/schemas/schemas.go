package schemas

type NumbersSetRequest struct {
	Numbers []int `json:"numbers"`
}

type NumbersSetResponse struct {
	Results map[string]int `json:"results"`
	Details string         `json:"details,omitempty"`
	//Done    bool           `json:"-"`
}
