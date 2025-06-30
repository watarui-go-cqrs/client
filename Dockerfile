FROM golang:1.24-alpine3.22
RUN apk update && apk add git curl alpine-sdk
RUN mkdir /go/src/client
WORKDIR /go/src/client
COPY . /go/src/client
RUN go mod download
EXPOSE 8081
CMD ["go", "run", "cmd/server/main.go"]
