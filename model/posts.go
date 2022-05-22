package model

import "time"

type Post struct {
	Id          int       `json:"id"`
	UserId      int       `json:"userId" binding:"required"`
	Likes       int       `json:"likes" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	ImageName   string    `json:"imageName" binding:"required"`
}

func NewPost(id int, userId int, likes int, description string, createdAt time.Time, updatedAt time.Time, imageName string) *Post {
	return &Post{
		Id:          id,
		UserId:      userId,
		Likes:       likes,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
		ImageName:   imageName,
	}
}

func (p *Post) GetShortDescription() string {
	if len(p.Description) <= 10 {
		return p.Description
	}
	return p.Description[:10] + "..."
}
