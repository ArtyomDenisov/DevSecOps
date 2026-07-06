# --- Этап Сборки (Builder) ---
FROM golang:1.21-alpine AS builder
WORKDIR /app

# 1. Копируем исходный код приложения и файлы модулей целиком
COPY . .

# 2. go mod tidy автоматически скачает нужные библиотеки и сам создаст go.sum прямо внутри контейнера, если его нет
RUN go mod tidy

# 3. Сборка статически скомпилированного бинарника без CGO
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# --- Этап Финального Образа ---
# а. Минимальный базовый образ (distroless или alpine)
FROM alpine:3.19

# c. Отсутствие лишних пакетов (скачиваем только обновления безопасности, если нужно)
RUN apk update && apk upgrade --no-cache

# b. Запуск от непривилегированного пользователя
# Создаем системную группу и пользователя appuser
RUN addgroup -S appgroup && adduser -S appuser -G appgroup
USER appuser

WORKDIR /home/appuser

# Копируем только бинарник из этапа сборки
COPY --from=builder /app/main .

# d. Закрытие ненужных портов (открываем только явный порт приложения)
EXPOSE 8080

CMD ["./main"]