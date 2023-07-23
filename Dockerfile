ARG GO_VERSION=1.18
FROM golang:${GO_VERSION}-alpine as builder

ARG MAIN_APP=./cmd/api
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -ldflags="-w -s" -o main ./cmd/api

FROM alpine
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
EXPOSE 1323
CMD ["/app/main"]