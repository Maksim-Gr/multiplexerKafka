FROM golang:1.21 as builder


WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -C cmd/consumer  -o ../../build/consumer -ldflags="-s -w"


FROM scratch
COPY --from=builder /app/build ./

CMD ["./consumer"]

