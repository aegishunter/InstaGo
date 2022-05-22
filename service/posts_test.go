package service

import (
	"instago/model"
	"instago/repository"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// AddPost(userId int, fileName string, description string) error
func TestPostService_AddPost(t *testing.T) {
	var postRepository = repository.NewPostsRepositoryMockInterface(t)
	var postService = PostsService{repository: postRepository}

	var (
		paramUserId      = 1
		paramFileName    = "lanaya.jpg"
		paramDescription = "test"
	)
	postRepository.Mock.On("AddPost", paramUserId, paramFileName, paramDescription).Return(nil)
	err := postService.AddPost(paramUserId, paramFileName, paramDescription)
	assert.Nil(t, err)
}

// 	UpdateDescription(description string, id int) error
func TestPostService_UpdateDescription(t *testing.T) {
	var postRepository = repository.NewPostsRepositoryMockInterface(t)
	var postService = PostsService{repository: postRepository}
	var (
		paramDescription = "abcdefghi"
		paramId          = 1
	)
	postRepository.Mock.On("UpdateDescription", paramDescription, paramId).Return(nil)
	err := postService.UpdateDescription(paramDescription, paramId)
	assert.Nil(t, err)
}

// 	GetPostDetailsById(id int, user_id int) (*model.Post, error)
func TestPostService_GetPostDetailsById(t *testing.T) {
	var postRepository = repository.NewPostsRepositoryMockInterface(t)
	var postService = PostsService{repository: postRepository}
	var (
		paramId     = 1
		paramUserId = 1
		post        = model.NewPost(
			paramId,
			paramUserId,
			0,
			"test",
			time.Now(),
			time.Now(),
			"test.jpg",
		)
	)
	postRepository.Mock.On("GetPostDetailsById", paramId, paramUserId).Return(post, nil)
	currPost, err := postService.GetPostDetailsById(paramId, paramUserId)
	assert.NotEmpty(t, currPost)
	assert.Nil(t, err)

	assert.Equal(t, currPost.Id, post.Id)
	assert.Equal(t, currPost.UserId, post.UserId)
	assert.Equal(t, currPost.Likes, post.Likes)
	assert.Equal(t, currPost.Description, post.Description)
	assert.Equal(t, currPost.CreatedAt, post.CreatedAt)
	assert.Equal(t, currPost.UpdatedAt, post.UpdatedAt)
	assert.Equal(t, currPost.ImageName, post.ImageName)
}

// 	GetCountTotalPosts() (int, error)
func TestPostService_GetPageCount(t *testing.T) {
	var testCase = []struct {
		postCount int
		limit     int
		expected  []int
	}{
		{1, 5, []int{1}},
		{10, 5, []int{1, 2}},
		{13, 5, []int{1, 2, 3}},
		{7, 5, []int{1, 2}},
		{32, 5, []int{1, 2, 3, 4, 5, 6, 7}},
		{20, 5, []int{1, 2, 3, 4}},
	}

	for _, val := range testCase {
		assert.Equal(t, val.expected, getPageCount(val.postCount, val.limit))
	}

}

func TestPostService_GetAllPosts(t *testing.T) {
	var postRepository = repository.NewPostsRepositoryMockInterface(t)
	var postService = PostsService{repository: postRepository}
	var (
		expectedPost = []model.Post{
			*model.NewPost(1, 1, 1, "test", time.Now(), time.Now(), "test.jpg"),
			*model.NewPost(2, 1, 2, "test2", time.Now(), time.Now(), "test2.jpg"),
		}
		paramIdx          = 1
		expectedPageCount = []int{1}
	)
	postRepository.Mock.On("GetAllPosts", paramIdx).Return(expectedPost, nil)
	postRepository.Mock.On("GetCountTotalPosts").Return(len(expectedPost), nil)
	allPosts, countTotalPosts, pageCount, err := postService.GetAllPosts(paramIdx)
	assert.NotNil(t, allPosts)
	assert.NotNil(t, countTotalPosts)
	assert.NotNil(t, pageCount)
	assert.Nil(t, err)

	for i := 0; i < len(allPosts); i++ {
		assert.Equal(t, expectedPost[i].Id, allPosts[i].Id)
		assert.Equal(t, expectedPost[i].UserId, allPosts[i].UserId)
		assert.Equal(t, expectedPost[i].Likes, allPosts[i].Likes)
		assert.Equal(t, expectedPost[i].Description, allPosts[i].Description)
		assert.Equal(t, expectedPost[i].CreatedAt, allPosts[i].CreatedAt)
		assert.Equal(t, expectedPost[i].UpdatedAt, allPosts[i].UpdatedAt)
		assert.Equal(t, expectedPost[i].ImageName, allPosts[i].ImageName)
	}

	assert.Equal(t, expectedPageCount, pageCount)
	assert.Equal(t, len(expectedPost), countTotalPosts)

}

// 	DeletePostById(id int, user_id int) error
func TestPostService_DeletePostById(t *testing.T) {
	var postRepository = repository.NewPostsRepositoryMockInterface(t)
	var postService = PostsService{repository: postRepository}
	var (
		paramId     = 1
		paramUserId = 1
	)
	postRepository.Mock.On("DeletePostById", paramId, paramUserId).Return(nil)
	err := postService.DeletePostById(paramId, paramUserId)
	assert.Nil(t, err)
}
