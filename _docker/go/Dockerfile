FROM golang:1.22-alpine

WORKDIR /app

# go.mod va go.sum ni olib kelamiz
COPY go.mod go.sum ./
RUN go mod download

# source kodlarni copy qilamiz
COPY ./app ./app

EXPOSE 8080

# Dev rejimda `go run`
CMD ["go", "run", "./app/main.go"]
