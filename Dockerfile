# Stage 1. Build the binary
FROM golang:1.11

# add a non-privileged user
RUN useradd -u 10001 myapp

RUN mkdir -p /go/src/github.com/rumyantseva/nsk
ADD . /go/src/github.com/rumyantseva/nsk
WORKDIR /go/src/github.com/rumyantseva/nsk

# build the binary with go build
RUN CGO_ENABLED=0 go build \
	-o bin/nsk github.com/rumyantseva/nsk/cmd/nsk

# Stage 2. Run the binary
FROM scratch

ENV PORT 8080

COPY --from=0 /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY --from=0 /etc/passwd /etc/passwd
USER myapp

COPY --from=0 /go/src/github.com/rumyantseva/nsk/bin/nsk /nsk
EXPOSE $PORT

CMD ["/nsk"]
