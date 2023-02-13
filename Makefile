default: dev

dev:
	air
	#go run github.com/tunght13488/realworld-go

build:
	go build github.com/tunght13488/realworld-go

run:
	./realworld-go

test:
	go test -v ./...

.PHONY: default dev build run test