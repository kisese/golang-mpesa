FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD ../go.mod .
ADD ../go.sum .

RUN go mod download
ADD .. .

RUN go install github.com/githubnemo/CompileDaemon@latest

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main