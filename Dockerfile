# syntax=docker/dockerfile:1
FROM golang:1.24 AS builder

WORKDIR /app
COPY go.mod main.go ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /goapp

FROM scratch
COPY --from=builder /goapp /goapp
CMD ["/goapp"]
