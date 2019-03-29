FROM golang:1.12.1

WORKDIR /go/src/checker
COPY ./checker/* ./

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["checker"]
