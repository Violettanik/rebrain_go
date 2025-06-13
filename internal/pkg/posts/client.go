package posts

import (
	"context"
	"net/http"
)

type Post struct {
	ID      int
	Title   string
	Author  string
}

type PostsClient interface {
	GetPosts(ctx context.Context, userID string) ([]Post, *http.Response, error)
}
