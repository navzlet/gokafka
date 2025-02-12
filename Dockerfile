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

# Опционально: открываем порт, если приложение слушает HTTP (если нет – можно убрать)
EXPOSE 8080

# Запускаем приложение
CMD ["./app"]
