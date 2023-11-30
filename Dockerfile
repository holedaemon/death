FROM golang:1.21.3 AS builder

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./ ./


RUN go build . -o death

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=builder /app/death /death
ENTRYPOINT [ "/death" ]