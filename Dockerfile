FROM golang:1.19.3-alpine as builder
COPY go.mod go.sum /go/src/gitlab.com/yevishev/restaurant-customer/
WORKDIR /go/src/gitlab.com/yevishev/restaurant-customer
RUN go mod download
COPY . /go/src/gitlab.com/yevishev/restaurant-customer
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/restaurant-customer gitlab.com/yevishev/restaurant-customer

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/gitlab.com/yevishev/restaurant-customer/build/restaurant-customer /usr/bin/restaurant-customer
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/restaurant-customerer"]