# Используем официальный образ Golang в качестве базового
FROM golang:1.23

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код проекта
COPY . .
COPY internal/config/.env /internal/config/.env
COPY .env .env

# Компилируем бинарник
RUN go build -o librarian_bot ./cmd/main.go

# Указываем команду для запуска контейнера
CMD ["./librarian_bot"]
