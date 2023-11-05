FROM golang:latest

ENV GO111MODULE=on

WORKDIR /app

COPY ./server .

RUN go install github.com/cosmtrek/air@latest
RUN go mod download

# RUN env GIT_TERMINAL_PROMPT=1 go get github.com/uxioandrade/go-sql-tutorial
# CMD ["go","run","main.go"]
CMD ["air", "-c", ".air.toml"]