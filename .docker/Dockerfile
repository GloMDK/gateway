# Этап, на котором выполняется сборка приложения
FROM golang:1.23-alpine as builder
COPY . .
RUN go build -o /main main.go
# Финальный этап, копируем собранное приложение
FROM alpine
COPY --from=builder main /bin/main
ENTRYPOINT ["/bin/main"]