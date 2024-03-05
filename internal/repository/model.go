package repository

type Goods struct {
	ID          int32  `json:"id"`
	ProjectID   int32  `json:"project_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Priority    int32  `json:"priority"`
	Removed     bool   `json:"removed"`
	Created_at  string `json:"created_at"`
}
