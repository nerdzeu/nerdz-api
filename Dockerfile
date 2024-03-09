FROM golang:1.22 as builder
MAINTAINER Paolo Galeone <nessuno@nerdz.eu>

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM gcr.io/distroless/base-debian12
COPY --from=builder /app/nerdz-api /

EXPOSE 8080
ENTRYPOINT ["/nerdz-api"]
