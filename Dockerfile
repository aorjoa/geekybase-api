FROM golang:1.10.3

WORKDIR /go/src/app
COPY . .

RUN go get -u github.com/golang/dep/...

RUN dep ensure
RUN go build -o punyim-api main.go

EXPOSE 1323

CMD ["./punyim-api"]