up:
	sudo docker compose up -d
down:
	sudo docker compose down
up-build:
	sudo docker compose up --build -d
delete-volumes:
	sudo docker compose down --volumes

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/main.go
	./app
run:
	go run cmd/main.go

