FROM golang:1.11.2 as builder
WORKDIR /go/src/app
ADD . .
RUN go get -d -v ./... \
&& CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:3.8 as app
RUN apk update && apk add ca-certificates \
&& rm -rf /var/cache/apk/*
WORKDIR /root
COPY --from=builder /go/src/app/app .

ENV TOKEN token
ENV CREDENTIALS /path/to/firestore/credentials

CMD ["./app"]