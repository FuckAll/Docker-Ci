FROM alpine:latest

RUN apk add --no-cache --update-cache bash curl ca-certificates make git go && \
  curl -Lo /etc/apk/keys/andyshinn.rsa.pub "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/andyshinn.rsa.pub" && \
  curl -Lo glibc-${GLIBC_PKG_VERSION}.apk "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-${GLIBC_PKG_VERSION}.apk" && \
  curl -Lo glibc-bin-${GLIBC_PKG_VERSION}.apk "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-bin-${GLIBC_PKG_VERSION}.apk" && \
  curl -Lo glibc-i18n-${GLIBC_PKG_VERSION}.apk "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-i18n-${GLIBC_PKG_VERSION}.apk" && \
  apk add glibc-${GLIBC_PKG_VERSION}.apk glibc-bin-${GLIBC_PKG_VERSION}.apk glibc-i18n-${GLIBC_PKG_VERSION}.apk && rm glibc-* && \
  curl -fsSL "http://dn-cdn-proxy.qbox.me/docker-1.9.1.tgz" -o docker-1.9.1.tgz && tar zxf docker-1.9.1.tgz -C / && rm docker-1.9.1.tgz && \
  curl -fsSL "http://dn-cdn-proxy.qbox.me/script.tar.gz" -o script.tar.gz && tar zxf script.tar.gz -C /usr/bin/ && rm script.tar.gz && \
  curl -fsSL "http://dn-cdn-proxy.qbox.me/protoc.tar.gz" -o protoc.tar.gz && tar zxf protoc.tar.gz -C /usr/bin/ && rm protoc.tar.gz

ENV GOPATH /gopath
ENV GOBIN /app
ENV PATH $PATH:$GOROOT/bin:$GOBIN


# go get -u
COPY golib.sh /
RUN /golib.sh

CMD["/bin/bash"]
