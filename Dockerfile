# two stages
#
# build
# - go image (compilers, full userspace, shell)
# - compiles a static binary
# - no cgo
# - uses caching properly
#
# second stage is prod
# - non-root user
# - scratch/alpine
# - copy over static binary
# - entrypoint for the static binary
# - expose/healthcheck type metatdata

ARG GO_VERSION=1.23

FROM --platform=$BUILDPLATFORM golang:$GO_VERSION AS build

WORKDIR /build

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
        go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN --network=none \
    --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH \
        go build -o /cmd main.go


FROM alpine:3.22

RUN apk --no-cache add ca-certificates

EXPOSE 8080

USER 1001:1001

COPY --from=build --chown=1001:1001 /cmd /

ENTRYPOINT ["/cmd"]
