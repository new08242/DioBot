FROM golang:1.11.5-alpine3.8

ENV DEP_VERSION="0.5.0"

RUN apk add --no-cache \
        ca-certificates \
        curl \
        git \
        make \
        openssl; \
    curl -L -s https://github.com/golang/dep/releases/download/v${DEP_VERSION}/dep-linux-amd64 -o /bin/dep; \
    chmod +x /bin/dep; \
    rm -rf /var/cache/apk/*; \
    rm -rf /tmp/*


ENV APP_PATH /go/src/dio-bot
WORKDIR ${APP_PATH}
RUN go get -u github.com/tockins/realize

RUN mkdir app
COPY . .
EXPOSE 9999
