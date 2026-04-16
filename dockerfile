# ---------- Stage 1: Build ----------
FROM golang:1.26-alpine AS builder

WORKDIR /app

# install git
RUN apk add --no-cache git

# copy everything
COPY . .

# tidy + download
RUN go mod tidy
RUN go mod download

# build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app cmd/api/main.go

# ---------- Stage 2 ----------
FROM alpine:3.19

WORKDIR /app

COPY --from=builder /app/app .

EXPOSE 8080

CMD ["./app"]