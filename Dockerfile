FROM golang:latest as BUILDER

RUN mkdir /webparser
COPY . /webparser
WORKDIR /webparser

ADD https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz /usr/local
RUN set -x && \
    apt update && \
    apt install -y xz-utils && \
    xz -d -c /usr/local/upx-3.96-amd64_linux.tar.xz | \
    tar -xOf - upx-3.96-amd64_linux/upx > /bin/upx && \
    chmod a+x /bin/upx && \
    go get -d -v . && \
    CGO_ENABLED=0 go build -o main . && \
    strip --strip-unneeded main && \
    upx main

FROM alpine
WORKDIR /root/
COPY --from=BUILDER /webparser/main .
COPY --from=BUILDER /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 80

CMD ["/root/main"]
