FROM golang:1.14

WORKDIR /go/src/app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app src/main.go

EXPOSE 8000

ENTRYPOINT ["./app"]