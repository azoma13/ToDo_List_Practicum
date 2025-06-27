FROM golang:1.23.0-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o to_do_list .

FROM alpine:latest

COPY --from=builder /app/to_do_list /usr/local/bin/to_do_list
COPY --from=builder /app/.env.example /.env
COPY web /web

EXPOSE 7540

CMD ["/usr/local/bin/to_do_list"]