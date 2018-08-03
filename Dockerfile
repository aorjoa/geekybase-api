FROM golang:1.10.3

WORKDIR /go/src/app
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure
RUN go build -o punyim-api main.go

EXPOSE 1323

CMD ["./punyim-api"]