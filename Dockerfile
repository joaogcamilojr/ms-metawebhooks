FROM golang:1.22.5 AS builder

WORKDIR /app

COPY ./cmd ./

RUN go build -o goapp

FROM gcr.io/distroless/base-debian12 AS release

COPY --from=builder /app/goapp /goapp

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/goapp" ]
