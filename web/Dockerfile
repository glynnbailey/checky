FROM golang:1.12.1

WORKDIR /go/src/web
COPY ./web/* ./

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 80

CMD ["web"]
