# BUILD container
FROM golang:1.10.0 as builder

WORKDIR /go/src/github.com/lukasjarosch/shipping-application/consignment-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep
RUN dep init && dep ensure
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# RUN container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/lukasjarosch/shipping-application/consignment-cli .
COPY --from=builder /go/src/github.com/lukasjarosch/shipping-application/consignment-cli/consignment.json .

CMD ["./consignment-cli"]




