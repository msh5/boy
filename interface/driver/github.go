package driver

import (
	"context"

	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type GitHubClient struct {
	client *githubv4.Client
}

func NewGitHubClient(accessToken string) *GitHubClient {
	httpClient := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	))

	return &GitHubClient{
		client: githubv4.NewClient(httpClient),
	}
}

type UserGistFiles struct {
	User struct {
		Gist struct {
			Files []struct {
				Name string
				Text string
			}
		} `graphql:"gist(name: $gist_name)"`
	} `graphql:"user(login: $user_id)"`
}

func (c *GitHubClient) QueryGistFiles(ctx context.Context, userID string, gistName string) (*UserGistFiles, error) {
	var obj UserGistFiles

	variables := map[string]interface{}{
		"user_id":   githubv4.String(userID),
		"gist_name": githubv4.String(gistName),
	}

	if err := c.client.Query(ctx, &obj, variables); err != nil {
		return nil, err
	}

	return &obj, nil
}

type RepositoryBlob struct {
	Repository struct {
		Object struct {
			Blob struct {
				Text string
			} `graphql:"... on Blob"`
		} `graphql:"object(expression: $object_expr)"`
	} `graphql:"repository(owner: $repo_owner, name: $repo_name)"`
}

func (c *GitHubClient) QueryRepositoryBlob(ctx context.Context, repoOwner string, repoName string, path string) (
	*RepositoryBlob, error) {
	var obj RepositoryBlob

	variables := map[string]interface{}{
		"repo_owner":  githubv4.String(repoOwner),
		"repo_name":   githubv4.String(repoName),
		"object_expr": githubv4.String("HEAD:" + path),
	}

	if err := c.client.Query(ctx, &obj, variables); err != nil {
		return nil, err
	}

	return &obj, nil
}
