FROM golang:1.13-alpine

RUN mkdir -p /backend
WORKDIR /backend

COPY . .

RUN go get -v
RUN go build

CMD["./pressx"]
