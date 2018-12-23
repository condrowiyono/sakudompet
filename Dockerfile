FROM golang:1.10-alpine3.7 as build

WORKDIR /build

ADD . .

RUN go build -o app api/main.go

FROM alpine:3.7

COPY --from=build /build/app /app

ENTRYPOINT ["/app"]