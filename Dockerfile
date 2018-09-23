FROM golang:1.11.0

WORKDIR /go/src/app
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN dep ensure
RUN go build -o geekybase-api main.go

EXPOSE 1323

CMD ["./geekybase-api"]
