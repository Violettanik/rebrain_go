package posts

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"go.uber.org/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPostCount_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockPostsClient(ctrl)
	ctx := context.Background()
	userID := "test-user"

	// Настраиваем ожидаемый вызов с возвратом 3 постов
	mockClient.EXPECT().
		GetPosts(ctx, userID).
		Return([]Post{
			{ID: 1, Title: "Post 1"},
			{ID: 2, Title: "Post 2"},
			{ID: 3, Title: "Post 3"},
		}, &http.Response{StatusCode: 200}, nil)

	count, err := PostCount(ctx, mockClient, userID)

	assert.NoError(t, err)
	assert.Equal(t, 3, count)
}

func TestPostCount_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockPostsClient(ctrl)
	ctx := context.Background()
	userID := "test-user"

	// Настраиваем ожидаемый вызов с ошибкой
	expectedErr := errors.New("API error")
	mockClient.EXPECT().
		GetPosts(ctx, userID).
		Return(nil, &http.Response{StatusCode: 500}, expectedErr)

	count, err := PostCount(ctx, mockClient, userID)

	assert.Error(t, err)
	assert.Equal(t, 0, count)
	assert.Equal(t, expectedErr, err)
}

func TestPostCount_Empty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockPostsClient(ctrl)
	ctx := context.Background()
	userID := "test-user"

	// Настраиваем ожидаемый вызов с пустым списком
	mockClient.EXPECT().
		GetPosts(ctx, userID).
		Return([]Post{}, &http.Response{StatusCode: 200}, nil)

	count, err := PostCount(ctx, mockClient, userID)

	assert.NoError(t, err)
	assert.Equal(t, 0, count)
}
