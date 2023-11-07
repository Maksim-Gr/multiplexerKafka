FROM golang:1.21-alpine3.18 as builder

ENV GO111MODULE=on CGO_ENABLED=0
RUN apk add upx

WORKDIR /app
COPY . .

RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-s -w"  -o build/consumer cmd/consumer/main.go
RUN upx -9 build/consumer

FROM scratch
COPY --from=builder /app/build ./

CMD ["./consumer"]

