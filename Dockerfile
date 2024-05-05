FROM golang:1.19.0

WORKDIR /usr/src/app

# Used to reload app when code is changed
RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy