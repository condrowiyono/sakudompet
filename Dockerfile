FROM lwolf/golang-glide:0.12.3

ENV APP_PATH=/go/src/github.com/condrowiyono/sakudompet

RUN mkdir -p $APP_PATH
WORKDIR $APP_PATH

COPY glide.yaml glide.yaml
COPY glide.lock glide.lock

RUN glide install -v

VOLUME ["/build"]

ADD . $APP_PATH
CMD GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /build/app