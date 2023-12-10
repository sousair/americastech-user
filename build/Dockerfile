FROM golang:1.21.3-alpine3.17 AS compiler

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/http-server ./cmd/...

#

FROM scratch

COPY --from=compiler /app/bin/http-server /app/bin/http-server

ENTRYPOINT ["/app/bin/http-server"]
