run:
	go run main.go

gen:
	protoc --go_out=./post --go_opt=paths=source_relative --go-grpc_out=require_unimplemented_servers=false:./post --go-grpc_opt=paths=source_relative ./post.proto

MIGRATION_NAME = default

build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o appctl

build-osx:
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o appctl

new-migration:
	$(eval version := $(shell date +%Y%m%d%H%M%S))
	touch assets/migrations/$(version)_$(MIGRATION_NAME).up.sql
	touch assets/migrations/$(version)_$(MIGRATION_NAME).down.sql

test:
	go test ./... -v -coverprofile=profile.out -p 1
	go tool cover -func=profile.out

lint:
	golangci-lint run