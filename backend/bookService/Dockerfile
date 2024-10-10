# Используйте официальный образ Golang в качестве стадии сборки
FROM golang:latest as build

# Установите рабочую директорию
WORKDIR /app

# Копируйте файлы модуля Go
COPY go.mod go.sum ./

# Скачайте зависимости модуля Go
RUN go mod download

# Копируйте все остальные файлы приложения, включая конфигурацию
COPY . .

# Соберите приложение Go со статической линковкой
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/main.go

# Используйте минимальный базовый образ для финальной стадии
FROM alpine:latest as run

# Копируйте исполняемый файл приложения из образа сборки
COPY --from=build /app/main main

# Копируем конфигурационный файл в финальное изображение
COPY --from=build /app/config_docker/config.yaml config/config.yaml

# Открываем порт приложения
EXPOSE 8080

# Запускаем приложение
CMD ["./main"]
