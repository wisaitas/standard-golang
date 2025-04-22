FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY ../../../go.mod ../../../go.sum ./

RUN go mod download && go mod verify

COPY ../../../cmd/standard-service/ ./cmd/standard-service
COPY ../../../internal/standard-service/ ./internal/standard-service
COPY ../../../pkg/ ./pkg

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/standard-service/main.go

RUN swag init -g cmd/standard-service/main.go -o internal/standard-service/docs

FROM scratch AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]