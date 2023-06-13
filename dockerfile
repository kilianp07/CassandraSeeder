From golang:1.20
LABEL org.opencontainers.image.source="https://github.com/kilianp07/cassandraseeder"
WORKDIR /go/src/app

COPY . .

RUN go build -o main main.go

CMD ["./main"]
