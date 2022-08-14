package services

import (
	"microservices-course/src/api/config"
	"microservices-course/src/api/domain/github"
	"microservices-course/src/api/domain/repositories"
	"microservices-course/src/api/providers/github_provider"
	"microservices-course/src/api/utils/errors"
	"strings"
)

type reposService struct {
}

type reposServiceInterface interface {
	CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError)
}

var RepositoryService reposServiceInterface

func init() {
	RepositoryService = &reposService{}
}

func (s *reposService) CreateRepo(req repositories.CreateRepoRequest) (*repositories.CreateRepoResponse, errors.ApiError) {
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		return nil, errors.NewBadRequestApiError("Invalid repo name")
	}

	ghReq := &github.CreateRepoRequest{
		Name:        req.Name,
		Description: "any",
		Private:     false,
	}

	res, err := github_provider.CreateRepo(config.GetGithubAccessToken(), *ghReq)
	if err != nil {
		return nil, errors.NewApiError(err.StatusCode, err.Message)
	}

	return &repositories.CreateRepoResponse{
		ID:       res.ID,
		Name:     res.Name,
		FullName: res.FullName,
	}, nil
}
