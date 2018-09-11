# Build
FROM golang:1.10-alpine as build-env

# ENV DEP_VERSION=0.5.0
ENV REPO=github.com/findkim/risky-wheels

# RUN apt-get update && apt-get install && \
#     curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mkdir -p $GOPATH/src/${REPO}
WORKDIR $GOPATH/src/${REPO}
COPY . .
# RUN dep ensure
RUN go build . && ls -al


# Execute
FROM alpine:3.8

ENV REPO=github.com/findkim/risky-wheels

# Update and add dumb-init to help with process exit
RUN apk add --no-cache ca-certificates apache2-utils dumb-init bash

WORKDIR /app

# Add static binary and public files
COPY --from=build-env /go/src/$REPO/risky-wheels .
COPY wheel wheel

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["/app/risky-wheels"]
