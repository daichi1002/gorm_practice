package service

import (
	"context"
	"test/usecase"

	pb "test/post"
)

type PostsService struct {
	usecase *usecase.PostUsecase
}

func (s *PostsService)ListPosts(context.Context, *pb.ListPostsRequest) (*pb.ListPostsResponse, error) {
	posts, err := s.usecase.ListPost()

	if err != nil {
		return nil, err
	}

	return &pb.ListPostsResponse{
		Posts: s.usecase.ConvertPostsDomainToPb(posts),
	}, nil
}