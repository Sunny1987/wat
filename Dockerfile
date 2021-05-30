FROM golang:latest

RUN mkdir /webparser

COPY . /webparser

WORKDIR /webparser

RUN go build -o main .

CMD ["/webparser/main"]