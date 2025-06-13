package posts

import (
	"context"
)

func PostCount(ctx context.Context, client PostsClient, userID string) (int, error) {
	posts, _, err := client.GetPosts(ctx, userID)
	if err != nil {
		return 0, err
	}
	return len(posts), nil
}
