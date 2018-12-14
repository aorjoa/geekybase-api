FROM golang:1.11.2

WORKDIR /go/src/app
COPY . .

ENV GO111MODULE=on
RUN go build -o geekybase-api .

EXPOSE 1323

CMD ["./geekybase-api"]
