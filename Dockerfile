FROM alpine:latest

ENV GLIBC_PKG_VERSION=2.23-r1
RUN apk add --no-cache --update-cache bash curl ca-certificates make git openssh go && \
  curl -Lo /etc/apk/keys/andyshinn.rsa.pub "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/andyshinn.rsa.pub" && \
  curl -Lo glibc-${GLIBC_PKG_VERSION}.apk "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-${GLIBC_PKG_VERSION}.apk" && \
  curl -Lo glibc-bin-${GLIBC_PKG_VERSION}.apk "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-bin-${GLIBC_PKG_VERSION}.apk" && \
  curl -Lo glibc-i18n-${GLIBC_PKG_VERSION}.apk "https://github.com/andyshinn/alpine-pkg-glibc/releases/download/${GLIBC_PKG_VERSION}/glibc-i18n-${GLIBC_PKG_VERSION}.apk" && \
  apk add glibc-${GLIBC_PKG_VERSION}.apk glibc-bin-${GLIBC_PKG_VERSION}.apk glibc-i18n-${GLIBC_PKG_VERSION}.apk && rm glibc-* && \
  curl -fsSL "https://get.docker.com/builds/Linux/x86_64/docker-1.9.1.tgz" -o docker-1.9.1.tgz && tar zxf docker-1.9.1.tgz -C / && rm docker-1.9.1.tgz 

ENV GOPATH /gopath
ENV GOBIN /app
ENV PATH $PATH:$GOROOT/bin:$GOBIN


# go get -u
COPY golib.sh 
COPY protoc /app/
RUN /golib.sh

CMD ["/bin/bash"]
