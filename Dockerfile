# 快取
FROM    golang:alpine AS build_base
RUN     apk add bash ca-certificates git gcc g++ libc-dev make build-base
ENV     RUN_PATH=/drs PROJ_PATH=/build
RUN     mkdir -p $RUN_PATH
WORKDIR $RUN_PATH
ENV     GO111MODULE=on
COPY    go.mod .
COPY    go.sum .
RUN     go mod download

# Build 專案;
FROM    golang:alpine AS builder
LABEL   maintainer="DAP平台中心"
USER    root
ADD     . ${PROJ_PATH}
WORKDIR ${PROJ_PATH}
RUN     make test build pack \
        && tar -zxf drs-v*.tar.gz -C ${RUN_PATH} \
        && rm -rf ${PROJ_PATH}

# 打包 Image;
FROM    alpine
LABEL   maintainer="DAP平台中心"
USER    root
ENV     RUN_PATH=/ops
RUN     mkdir -p $RUN_PATH && apk add --no-cache ca-certificates bash git
COPY    --from=builder ${RUN_PATH} ${RUN_PATH}
WORKDIR ${RUN_PATH}
ENTRYPOINT ["./drs"]
