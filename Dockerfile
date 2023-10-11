FROM golang:1.17.9

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download all

COPY . ./

CMD ["go", "run", "/app/main.go"]