package main

import (
	"test/db"
	"test/infra/repository"
	"test/usecase"
)

func main() {
	database, _ := db.ConnectDatabaseWithGorm()
	postRepository := repository.NewPostRepository(database)

	postUsecase := usecase.NewPostUsecase(*postRepository)
}