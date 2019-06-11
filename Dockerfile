FROM golang:1.12.5-alpine3.9 AS builder

RUN apk update && apk upgrade && apk add --no-cache bash git openssh

WORKDIR /go/src/github.com/alextanhongpin/go-locale

COPY . .

RUN go get -v ./...

RUN CGO=0 GOOS=linux go build -o app .


FROM alpine:3.9

COPY --from=builder /go/src/github.com/alextanhongpin/go-locale/app .
# Copy all the locale files to the destination too.
COPY --from=builder /go/src/github.com/alextanhongpin/go-locale/*.toml /

CMD ["./app"]
