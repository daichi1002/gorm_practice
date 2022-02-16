package main

import (
	"fmt"
	"test/db"
	"test/infra/repository"
	"test/usecase"
)

func main() {
	database, _ := db.ConnectDatabaseWithGorm()
	postRepository := repository.NewPostRepository(database)

	postUsecase := usecase.NewPostUsecase(*postRepository)

	//posts一覧取得
	posts, err := postUsecase.ListPost()
	fmt.Println(posts, err)
}