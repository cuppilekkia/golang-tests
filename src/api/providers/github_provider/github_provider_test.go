package github_provider

import (
	"microservices-course/src/api/domain/github"
	"reflect"
	"testing"
)

func Test_getAuthHeader(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"Make header OK", args{"1234"}, "token 1234"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAuthHeader(tt.args.t); got != tt.want {
				t.Errorf("getAuthHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateRepo(t *testing.T) {
	type args struct {
		accessToken string
		req         github.CreateRepoRequest
	}
	tests := []struct {
		name  string
		args  args
		want  *github.CreateRepoResponse
		want1 *github.GithubErrorResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := CreateRepo(tt.args.accessToken, tt.args.req)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateRepo() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CreateRepo() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
