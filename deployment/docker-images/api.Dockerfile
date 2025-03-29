FROM golang:1.23.2-alpine

WORKDIR /app

COPY ../../go.mod ../../go.sum ./

RUN go mod download && go mod verify

COPY ../../cmd ./cmd
COPY ../../data ./data
COPY ../../internal ./internal
COPY ../../pkg ./pkg
RUN go mod tidy

RUN go build -o main cmd/main.go

RUN chmod +x main

CMD ["./main"]