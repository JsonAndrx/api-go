FROM golang:1.22.1-bullseye

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main cmd/main.go

WORKDIR /app

RUN mkdir "/builld"

COPY . .

RUN go get github.com/githubnemo/CompileDaemon

RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --polling=true --build="go build -o /app/main ./cmd/main.go" --command="/app/main"