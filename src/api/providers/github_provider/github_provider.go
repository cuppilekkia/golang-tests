package github_provider

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"microservices-course/src/api/clients/restclient"
	"microservices-course/src/api/domain/github"
	"net/http"
)

const (
	headerAuth       = "Authorization"
	headerAuthFormat = "token %s"

	urlCreateRepo = "https://api.github.com/user/repos"
)

func getAuthHeader(t string) string {
	return fmt.Sprintf(headerAuthFormat, t)
}

func CreateRepo(accessToken string, req github.CreateRepoRequest) (*github.CreateRepoResponse, *github.GithubErrorResponse) {
	headers := http.Header{}
	headers.Set(headerAuth, getAuthHeader(accessToken))

	response, err := restclient.Post(urlCreateRepo, req, headers)

	if err != nil {
		return nil, &github.GithubErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	bytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, &github.GithubErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	defer response.Body.Close()

	if response.StatusCode > 299 {
		var errRes github.GithubErrorResponse
		if err := json.Unmarshal(bytes, &errRes); err != nil {
			return nil, &github.GithubErrorResponse{
				Message:    err.Error(),
				StatusCode: http.StatusInternalServerError,
			}
		}
		errRes.StatusCode = response.StatusCode
		return nil, &errRes
	}

	var result github.CreateRepoResponse

	if err := json.Unmarshal(bytes, &result); err != nil {
		return nil, &github.GithubErrorResponse{
			Message:    err.Error(),
			StatusCode: http.StatusInternalServerError,
		}
	}

	return &result, nil
}
