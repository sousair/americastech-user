FROM golang:1.21.3-alpine3.17 AS compiler

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin ./cmd/...

#

FROM scratch

COPY --from=compiler /app/bin /app/bin

COPY .env .env

ENTRYPOINT ["/app/bin/httpserver"]
