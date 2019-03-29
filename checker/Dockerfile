FROM golang:1.12.1

WORKDIR /go/src/github.com/glynnbailey/checky
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
