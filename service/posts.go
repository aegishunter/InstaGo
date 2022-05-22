package service

import (
	"instago/model"
	"instago/repository"
)

type PostsService struct {
	repository repository.PostsRepositoryInterface
}

func NewPostsService() *PostsService {
	return &PostsService{
		repository: repository.NewPostsRepository(),
	}
}

func getPageCount(totalPostCount int, limit int) []int {
	var n int
	if totalPostCount%limit == 0 {
		n = totalPostCount / limit
	} else {
		n = totalPostCount/limit + 1
	}
	res := make([]int, n)
	for i := range res {
		res[i] = i + 1
	}
	return res
}

func (p *PostsService) AddPost(userId int, fileName string, description string) error {
	return p.repository.AddPost(userId, fileName, description)
}

func (p *PostsService) UpdateDescription(description string, id int) error {
	return p.repository.UpdateDescription(description, id)
}

func (p *PostsService) GetAllPosts(idx int) ([]model.Post, int, []int, error) {
	allPosts, err := p.repository.GetAllPosts(idx)
	if err != nil {
		return nil, -1, nil, err
	}
	countTotalPosts, err1 := p.repository.GetCountTotalPosts()
	if err != nil {
		return nil, -1, nil, err1
	}
	return allPosts, countTotalPosts, getPageCount(countTotalPosts, 5), nil
}

func (p *PostsService) GetPostDetailsById(id int, user_id int) (*model.Post, error) {
	return p.repository.GetPostDetailsById(id, user_id)
}

func (p *PostsService) DeletePostById(id int, user_id int) error {
	return p.repository.DeletePostById(id, user_id)
}
