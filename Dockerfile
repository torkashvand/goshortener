# Build Stage
FROM lacion/alpine-golang-buildimage:1.14 AS build-stage

LABEL app="build-goshortener"
LABEL REPO="https://github.com/torkashvand/goshortener"

ENV PROJPATH=/go/src/github.com/torkashvand/goshortener

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/torkashvand/goshortener
WORKDIR /go/src/github.com/torkashvand/goshortener

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/torkashvand/goshortener"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/goshortener/bin

WORKDIR /opt/goshortener/bin

COPY --from=build-stage /go/src/github.com/torkashvand/goshortener/bin/goshortener /opt/goshortener/bin/
RUN chmod +x /opt/goshortener/bin/goshortener

# Create appuser
RUN adduser -D -g '' goshortener
USER goshortener

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/goshortener/bin/goshortener"] 
