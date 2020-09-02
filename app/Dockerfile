FROM golang:1.14 as builder

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app src/main.go

FROM alpine:3.10

WORKDIR /app

COPY --from=builder /go/src/app .

EXPOSE 8000

ENTRYPOINT ["./app"]