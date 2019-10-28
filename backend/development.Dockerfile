FROM golang:1.13-alpine

RUN mkdir -p /backend
WORKDIR /backend

COPY . .

RUN go get -v; go get github.com/cortesi/modd/cmd/modd
RUN go build

CMD ["./pressx"]
