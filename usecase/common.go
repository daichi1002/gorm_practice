package usecase

import (
	"test/domain/model"
	pb "test/post"
)

func convertPostDomainToPb(post *model.Post) *pb.Post {
	return &pb.Post{
		Id: post.Id,
		Title: post.Title,
		Content: post.Content,
	}
}