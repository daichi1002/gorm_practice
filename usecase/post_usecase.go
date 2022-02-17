package usecase

import (
	"test/domain/model"
	"test/infra/repository"
	pb "test/post"
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

func (u *PostUsecase) ConvertPostsDomainToPb(posts []*model.Post) []*pb.Post {
	var pbPosts  []*pb.Post
	for _, post := range posts {
		pbPosts = append(pbPosts, convertPostDomainToPb(post))
	}

	return pbPosts
}