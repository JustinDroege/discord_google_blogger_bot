package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/JustinDroege/BloggerBot/pkg/models"
)

type BloggerAPI struct {
	baseUrl  string
	client   *http.Client
	apiToken string
	blogID   string
}

func NewBloggerAPI(apiToken string, blogID string, client *http.Client, baseUrl string) (*BloggerAPI, error) {
	if apiToken == "" {
		return nil, fmt.Errorf("apiToken is required")
	}
	if blogID == "" {
		return nil, fmt.Errorf("blogID is required")
	}
	if len(baseUrl) == 0 {
		baseUrl = "https://www.googleapis.com/blogger/v3/"
	}

	return &BloggerAPI{
		client:   client,
		apiToken: apiToken,
		blogID:   blogID,
		baseUrl:  baseUrl,
	}, nil
}

func (b *BloggerAPI) GetPosts(pageToken string, date string) (*models.Posts, error) {
	url := fmt.Sprintf("%sblogs/%s/posts?key=%s", b.baseUrl, b.blogID, b.apiToken)

	if pageToken != "" {
		url = fmt.Sprintf("%s&pageToken=%s", url, pageToken)
	}

	if date != "" {
		url = fmt.Sprintf("%s&startDate=%s", url, date)
	}

	resp, err := b.client.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		return nil, fmt.Errorf("Error: %s", resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var posts models.Posts
	err = json.Unmarshal(body, &posts)

	if err != nil {
		return nil, err
	}

	return &posts, nil
}
