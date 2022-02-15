package model

import "time"

type Post struct {
	Id uint64
	Title string
	Content string
	CreatedAt time.Time
	UpdatedAt time.Time
}