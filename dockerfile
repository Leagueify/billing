# download go depedencies
FROM golang:1.23.0-alpine3.20 AS base
WORKDIR /app
COPY go.mod ./
RUN go mod download
# build local development image
FROM golang:1.23.0-alpine3.20 AS dev
COPY --from=base /go/bin /go/bin
WORKDIR /app
RUN go install github.com/air-verse/air@latest
EXPOSE 6507
# build the go binary
FROM golang:1.23.0-alpine3.20 AS builder
COPY --from=base /go/bin /go/bin
WORKDIR /app
COPY . ./
EXPOSE 6507
RUN CGO_ENABLED=0 GOOS=linux go build -o bin/billing cmd/billing.go
# create production image
FROM gcr.io/distroless/base-debian11 AS release
COPY --from=builder /app/bin/billing /billing
EXPOSE 6507
ENTRYPOINT ["/billing"]
