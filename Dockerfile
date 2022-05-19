##
## Build
##
FROM golang:alpine as builder

WORKDIR /build

ADD *.go /build
RUN go mod init build

RUN go build -o /webapp

##
## Deploy
##
FROM alpine
MAINTAINER abowspy<abowspy@gmail.com>

WORKDIR /

COPY --from=builder /webapp /webapp

EXPOSE 80

CMD ["/webapp"]