package repository

import (
	"instago/model"
)

type PostsRepository struct {
}

type PostsRepositoryInterface interface {
	AddPost(userId int, fileName string, description string) error
	UpdateDescription(description string, id int) error
	GetPostDetailsById(id int, user_id int) (*model.Post, error)
	GetCountTotalPosts() (int, error)
	GetAllPosts(idx int) ([]model.Post, error)
	DeletePostById(id int, user_id int) error
}

func NewPostsRepository() *PostsRepository {
	return &PostsRepository{}
}

func (p *PostsRepository) AddPost(userId int, fileName string, description string) error {
	sql := "insert into posts(\"userId\", likes, description, \"imageName\") values ($1,0,$2,$3)"
	db := getInstance()
	_, err := db.Exec(sql, userId, description, fileName)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostsRepository) UpdateDescription(description string, id int) error {
	sql := "UPDATE posts SET description = $1, \"updatedAt\" = now() WHERE id = $2"
	db := getInstance()
	_, err := db.Exec(sql, description, id)
	if err != nil {
		return err
	}
	return nil
}

func (p *PostsRepository) GetPostDetailsById(id int, user_id int) (*model.Post, error) {
	var selectedPost model.Post
	db := getInstance()
	sql := "SELECT * FROM posts WHERE id = $1 AND \"userId\" = $2 LIMIT 1"
	data, err := db.Query(sql, id, user_id)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		err := data.Scan(&selectedPost.Id, &selectedPost.UserId, &selectedPost.Likes, &selectedPost.Description,
			&selectedPost.CreatedAt, &selectedPost.UpdatedAt, &selectedPost.ImageName)
		if err != nil {
			panic(err)
		}
	}
	return &selectedPost, nil
}

func (p *PostsRepository) GetCountTotalPosts() (int, error) {
	db := getInstance()
	sql := "SELECT COUNT(1) FROM posts"
	data, err := db.Query(sql)
	if err != nil {
		return -1, err
	}
	defer data.Close()
	var count int

	for data.Next() {
		err = data.Scan(&count)
		if err != nil {
			return -1, nil
		}
	}
	return count, nil
}

func (p *PostsRepository) GetAllPosts(idx int) ([]model.Post, error) {
	var allPosts []model.Post
	limit := 5
	offset := limit * (idx - 1)
	db := getInstance()
	sql := "SELECT * FROM posts ORDER BY \"updatedAt\" DESC LIMIT $1 offset $2"
	data, err := db.Query(sql, limit, offset)
	if err != nil {
		return nil, err
	}
	defer data.Close()
	for data.Next() {
		var newPost model.Post
		err := data.Scan(&newPost.Id, &newPost.UserId, &newPost.Likes, &newPost.Description, &newPost.CreatedAt,
			&newPost.UpdatedAt, &newPost.ImageName)
		if err != nil {
			return nil, err
		}
		allPosts = append(allPosts, newPost)
	}
	return allPosts, nil
}

func (p *PostsRepository) DeletePostById(id int, user_id int) error {
	sql := "DELETE FROM posts WHERE id = $1 and \"userId\" = $2"
	db := getInstance()
	_, err := db.Exec(sql, id, user_id)
	if err != nil {
		return err
	}
	return nil
}
