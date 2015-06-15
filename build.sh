#!/bin/sh

# Environment list
# $GOOS     $GOARCH
# darwin    386
# darwin    amd64
# freebsd   386
# freebsd   amd64
# freebsd   arm
# linux     386
# linux     amd64
# linux     arm
# netbsd    386
# netbsd    amd64
# netbsd    arm
# openbsd   386
# openbsd   amd64
# plan9     386
# plan9     amd64
# windows   386
# windows   amd64

set -e

OS=("darwin" "darwin" "freebsd" "freebsd" "freebsd" "linux" \
  "linux" "linux" "netbsd" "netbsd" "netbsd" "openbsd" "openbsd" \
  "plan9" "plan9" "windows" "windows")
ARCH=("386" "amd64" "386" "amd64" "arm" "386" "amd64" "arm" \
  "386" "amd64" "arm" "386" "amd64" "386" "amd64" "386" "amd64")

mkdir -p ./dist
for i in `seq 0 1 16`
do
  GOOS=${OS["$i"]}
  GOARCH=${ARCH["$i"]}
  echo "rssapi.${GOOS}.${GOARCH}"
  GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./dist/rssapi.${GOOS}.${GOARCH} ./*.go
done
