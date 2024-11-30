FROM golang:1.23.2 AS builder

WORKDIR /app

COPY ./ ./

WORKDIR /app/cmd

RUN go mod download

RUN go build -o goapp

FROM gcr.io/distroless/base-debian12 AS release

COPY --from=builder /app/cmd/goapp /goapp

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/goapp" ]
