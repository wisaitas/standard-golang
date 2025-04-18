FROM golang:1.23.2-alpine AS builder

WORKDIR /app

COPY ../../../go.mod ../../../go.sum ./

RUN go mod download && go mod verify

COPY ../../../cmd ./cmd
COPY ../../../internal ./internal
COPY ../../../pkg ./pkg

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

FROM scratch AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]