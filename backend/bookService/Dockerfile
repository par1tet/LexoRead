FROM golang:latest as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

FROM alpine:latest as run

COPY --from=build /app/main main
COPY --from=build /app/config_docker/config.yaml config/config.yaml

CMD ["./main"]
