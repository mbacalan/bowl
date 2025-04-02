FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o bowl .

FROM alpine:3

WORKDIR /app

RUN apk --no-cache add ca-certificates sqlite

COPY --from=builder /app/.env .
COPY --from=builder /app/bowl .

EXPOSE 3000

CMD ["./bowl"]
