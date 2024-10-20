FROM golang:1.23

WORKDIR /go/src/source

COPY . . 


RUN go get -v ./...

RUN go build -o main .

CMD ["./main"]