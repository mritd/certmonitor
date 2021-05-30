FROM golang:1.16-alpine AS builder

ENV SRC_PATH ${GOPATH}/src/github.com/mritd/certmonitor

COPY . ${SRC_PATH}

WORKDIR ${SRC_PATH}

RUN set -ex \
    && export BUILD_VERSION=$(cat version) \
    && export BUILD_DATE=$(date "+%F %T") \
    && export COMMIT_SHA1=$(git rev-parse HEAD) \
    && go install -ldflags \
        "-X 'github.com/mritd/certmonitor/cmd.version=${BUILD_VERSION}' \
        -X 'github.com/mritd/certmonitor/cmd.buildDate=${BUILD_DATE}' \
        -X 'github.com/mritd/certmonitor/cmd.commit=${COMMIT_SHA1}'" \
    && scanelf --needed --nobanner /go/bin/certmonitor | \
        awk '{ gsub(/,/, "\nso:", $2); print "so:" $2 }' | \
        sort -u | tee /dep_so


FROM alpine:3.12

LABEL maintainer="mritd <mritd@linux.com>"

ARG TZ="Asia/Shanghai"

ENV TZ ${TZ}
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV LANGUAGE en_US:en

# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

RUN apk upgrade \
    && apk add bash tzdata ca-certificates \
    && ln -sf /usr/share/zoneinfo/${TZ} /etc/localtime \
    && echo ${TZ} > /etc/timezone \
    && rm -rf /var/cache/apk/*

COPY --from=builder /dep_so /dep_so
COPY --from=builder /go/bin/certmonitor /usr/bin/poetbot

CMD ["certmonitor"]