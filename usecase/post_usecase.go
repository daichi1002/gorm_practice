package usecase

import (
	"test/domain/model"
	"test/infra/repository"
)

type PostUsecase struct {
	repository repository.PostRepository
}

func NewPostUsecase(postRepository repository.PostRepository) *PostUsecase {
	return &PostUsecase{repository: postRepository}
}

func (u *PostUsecase) ListPost() ([]*model.Post, error) {
	posts, err := u.repository.ListPosts()
	return posts, err
}

func (u *PostUsecase) CreatePost(title string, content string) (*model.Post, error) {
	post, err := u.repository.CreatePost(title, content)
	return post, err
}