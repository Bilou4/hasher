#!/usr/bin/env bash

set -e

GO_FLAGS=(-ldflags "-w -s" -trimpath)
ENABLE_CGO=0
PACKAGE_NAME=hasher
OUTPUT_DIR=build

PLATFORMS=("windows/amd64" "windows/386" "linux/amd64" "linux/386")


for platform in "${PLATFORMS[@]}"
do
    # split platform at the '/' character and save the array value in platform_split
    IFS='/' read -ra platform_split <<< "$platform"
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    OUTPUT_NAME=$OUTPUT_DIR/$PACKAGE_NAME'-'$GOOS'-'$GOARCH
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi
    echo [+] Building $PACKAGE_NAME for "$GOOS" "$GOARCH"
    env CGO_ENABLED="$ENABLE_CGO" GOOS="$GOOS" GOARCH="$GOARCH" go build "${GO_FLAGS[@]}" -o "$OUTPUT_NAME" .
done