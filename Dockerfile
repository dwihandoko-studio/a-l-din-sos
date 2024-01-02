FROM golang:1.17.9

WORKDIR /app

#COPY go.mod ./
#COPY go.sum ./

#RUN go mod download all

COPY . ./
RUN go mod tidy

# Build the Go app
RUN go build -o main .

# Define a command to run the executable
CMD ["./main"]
#CMD ["go", "run", "/app/main.go"]

# FROM golang:1.17.9

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download all

# COPY . ./

# CMD ["go", "run", "/app/main.go"]