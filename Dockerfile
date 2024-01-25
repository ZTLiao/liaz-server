FROM golang:1.21.3

USER root

RUN apt-get update && \
    apt-get install -y make && \
    rm -rf /var/lib/apt/lists/*

ARG PROFILES_ACTIVE
ARG APPLICATION_NAME
ARG SERVER_PORT

ENV WORK_DIR /data/golang/$APPLICATION_NAME
ENV PROFILES_ACTIVE $PROFILES_ACTIVE
ENV SERVER_PORT $SERVER_PORT 
ENV GO111MODULE=on \
	GOPATH=$WORK_DIR/go \
	GOPROXY=https://goproxy.cn,direct

WORKDIR $WORK_DIR/

RUN mkdir -p $WORK_DIR/logs

CMD ["make", "param1"]

HEALTHCHECK --interval=20s --timeout=10s --retries=10 CMD wget --quiet --tries=1 --spider http://localhost:$SERVER_PORT || exit 1

