package repositories

type CreateRepoRequest struct {
	Name string `json:"name"`
}

type CreateRepoResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	FullName string `json:"fullName"`
}
