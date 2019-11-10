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

type UserGistQuery struct {
	User struct {
		Gist struct {
			Files []struct {
				Name string
				Text string
			}
		} `graphql:"gist(name: $name)"`
	} `graphql:"user(login: $user)"`
}

func (c *GitHubClient) QueryUserGist(ctx context.Context, userID string, gistName string) (*UserGistQuery, error) {
	var q UserGistQuery

	variables := map[string]interface{}{
		"user": githubv4.String(userID),
		"name": githubv4.String(gistName),
	}

	if err := c.client.Query(ctx, &q, variables); err != nil {
		return nil, err
	}

	return &q, nil
}
