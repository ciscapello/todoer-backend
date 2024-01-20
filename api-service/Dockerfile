FROM golang:1.21.4-alpine3.18

RUN mkdir /app
EXPOSE 3000
WORKDIR /app
ADD . /app

RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build cmd/main.go" --command="./main"
