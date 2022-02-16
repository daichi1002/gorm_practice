package repository

import (
	"test/domain/model"

	"gorm.io/gorm"
)

type PostRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) ListPosts() ([]*model.Post, error) {
	posts := make([]*model.Post, 0)
	err := r.db.Find(&posts).Error
	return posts, err
}