package mugshot_go

type AddFaceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	FaceID  string `json:"face_id"`
}

type SearchFaceItem struct {
	ID       string                 `json:"id"`
	Score    float64                `json:"score"`
	Metadata map[string]interface{} `json:"metadata"`
}

type SearchFaceResponse struct {
	Success bool             `json:"success"`
	Result  []SearchFaceItem `json:"result"`
}

type MatchFaceItem struct {
	ID       string                 `json:"id"`
	Match    bool                   `json:"match"`
	Score    float64                `json:"score"`
	Metadata map[string]interface{} `json:"metadata"`
}

type MatchFaceResponse struct {
	Success bool            `json:"success"`
	Result  []MatchFaceItem `json:"result"`
}

type DeleteFaceResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
