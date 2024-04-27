FROM golang:1.22-bookworm AS builder

WORKDIR /app

COPY . .

RUN go build -o mock-shipping-provider .

FROM debian:bookworm-slim AS runtime

WORKDIR /app

RUN apt-get update && apt-get install -y curl ca-certificates sqlite3

COPY --from=builder /app/mock-shipping-provider .

ENV HTTP_HOSTNAME="0.0.0.0"
ENV HTTP_PORT="3000"

EXPOSE ${HTTP_PORT}

CMD ["./mock-shipping-provider"]
