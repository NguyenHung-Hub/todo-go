# FROM golang:1.21.1-alpine3.17

# RUN apk add git
# WORKDIR /app
# COPY . .

# ENV CGO_ENABLED=0

# RUN go get github.com/githubnemo/CompileDaemon

# CMD CompileDaemon -build="go build" -command="./main"

FROM golang:1.21.1-alpine
WORKDIR /app
RUN apk update && \
    apk add libc-dev && \
    apk add gcc && \
    apk add git

ENV PROJECT_DIR=/app \
    GO111MODULE=on \
    CGO_ENABLED=0

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

RUN go install github.com/githubnemo/CompileDaemon@latest
COPY . .

ENTRYPOINT CompileDaemon --build="go build -o main main.go"  --command=./main

# RUN go build -o /go-docker

# CMD [ "/go-docker" ]