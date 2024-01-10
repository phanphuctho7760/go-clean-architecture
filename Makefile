run:
	go run ./app/entry/api/main.go
start:
	go run ./app/entry/api/main.go
watch:
	air -d
build:
	go build -o ./app/entry/api ./app/entry/api
mod:
	go mod tidy && go mod vendor
test:
	go test ./... -v -cover
swag:
	swag fmt && swag init -g ./app/entry/api/main.go