package usecase

import "test/infra/repository"

type PostUsecase struct {
	repository repository.PostRepository
}

func NewPostUsecase(postRepository repository.PostRepository) *PostUsecase {
	return &PostUsecase{repository: postRepository}
}