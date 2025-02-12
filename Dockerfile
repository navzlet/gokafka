# Используем официальный образ Go
FROM golang:1.23.6-alpine

WORKDIR /app

# Копируем файлы зависимостей и скачиваем их
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь исходный код
COPY . .

# Собираем бинарник (учтите, что main находится в cmd/main.go)
RUN go build -o app cmd/main.go

# Запускаем приложение
CMD ["./app"]
