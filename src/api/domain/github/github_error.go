package github

type GithubErrorResponse struct {
	Message          string        `json:"message"`
	Errors           []GithubError `json:"errors"`
	DocumentationURL string        `json:"documentation_url"`
	StatusCode       int           `json:"status_code"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
