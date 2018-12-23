FROM golang:1.10-alpine3.7 as build

WORKDIR /go/src/github.com/condrowiyono/sakudompet

COPY . .

RUN go build -o app api/main.go

FROM alpine:3.7

COPY --from=build /go/src/github.com/condrowiyono/sakudompet/app /usr/local/bin/app

ENTRYPOINT ["/usr/local/bin/app"]