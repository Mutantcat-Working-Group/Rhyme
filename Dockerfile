# Rhyme 容器镜像：多阶段构建，产出 CGO_ENABLED=0 静态二进制。
# 挂载诗歌文件或目录后调用，例如：
#   docker run --rm -v "$PWD/poems":/poems ghcr.io/mutantcat-working-group/rhyme:latest -folder=/poems -key=风
FROM golang:1.23-alpine AS build

ARG VERSION=1.0.20260920

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -trimpath \
    -ldflags "-s -w -X main.version=${VERSION}" \
    -o /out/rhyme .

FROM alpine:3.21

RUN adduser -D -H -u 10001 rhyme \
    && mkdir -p /work \
    && chown rhyme:rhyme /work

COPY --from=build /out/rhyme /usr/local/bin/rhyme

USER rhyme
WORKDIR /work

ENTRYPOINT ["/usr/local/bin/rhyme"]
