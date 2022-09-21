FROM golang:1.19.0-bullseye AS builder

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

WORKDIR /app

COPY src/ ./

RUN go mod download
RUN go mod verify
RUN go build -o /auth

FROM golang:1.19.0-bullseye

WORKDIR /

COPY --from=builder /auth /auth

EXPOSE 8082
CMD ["/auth"]