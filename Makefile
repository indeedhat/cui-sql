default: build

cover:
	@go test -coverprofile=cover.out
	@go tool cover -html=cover.out
	@rm cover.out

test:
	@go test -cover

build:
	@CGO_ENABLED=0 go build

run:
	@CGO_ENABLED=0 go build
	@./cui-sql
