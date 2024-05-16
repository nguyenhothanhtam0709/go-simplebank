# Build stage
FROM golang:1.22.2-alpine3.19 AS builder 
WORKDIR /app

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz

COPY . .
RUN go build -o main main.go

# Run stage
FROM alpine:3.19
WORKDIR /app

COPY start.sh wait-for.sh ./
COPY --from=builder /app/migrate ./migrate

COPY db/migrations ./migrations
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 8080
# CMD provide default argument for ENTRYPOINT
CMD [ "/app/main" ] 
ENTRYPOINT [ "/app/start.sh" ]