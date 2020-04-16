FROM golang:1.13-alpine AS dev

RUN apk add --no-cache git

RUN env GO111MODULE=on go get -v github.com/cortesi/modd/cmd/modd

COPY . /pressx

WORKDIR /pressx

RUN go get -v && go build -v

CMD ["/pressx/pressx"]

FROM alpine:latest AS build

COPY --from=dev /pressx/pressx /pressx

CMD ["/pressx"]
