FROM golang:latest as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build server.go

RUN CGO_ENABLED=0 GOOS=linux go build server.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/server .
COPY --from=builder /app/.env .

EXPOSE 8081

CMD ["./server"]