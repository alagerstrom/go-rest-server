FROM golang:1.8

WORKDIR /go/src/app
COPY . .

EXPOSE 1337

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]