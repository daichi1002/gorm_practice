package main

import (
	"test/db"
	"test/infra/repository"
)

func main() {
	database, _ := db.ConnectDatabaseWithGorm()
	postRepository := repository.NewPostRepository(database)

}