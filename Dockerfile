FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# Final stage
FROM scratch

WORKDIR /root/

COPY --from=builder /app/main .

EXPOSE 3002

CMD ["./main"]