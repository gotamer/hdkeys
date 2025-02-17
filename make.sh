#!/usr/bin/env bash

set -e

APPNAME='hdkeys'
PACKAGE='github.com/gotamer/hdkeys'

clear

cd $( dirname -- "$0"; );
ROOTDIR=$PWD;

VERSION="$(git describe --tags --always --abbrev=0 --match='v[0-9]*.[0-9]*.[0-9]*' 2> /dev/null | sed 's/^.//')"
COMMIT_HASH="$(git rev-parse --short HEAD)"
BUILD_TIME=$(date '+%Y-%m-%dT%H:%M:%S')
BUILD_USER=$(id -u -n)

LDFLAGS=("-s -w "
  "-X '${PACKAGE}/build.AppName=${APPNAME}'"
  "-X '${PACKAGE}/build.Version=${VERSION}'"
  "-X '${PACKAGE}/build.CommitHash=${COMMIT_HASH}'"
  "-X '${PACKAGE}/build.BuildTime=${BUILD_TIME}'"
  "-X '${PACKAGE}/build.UserName=${BUILD_USER}'"
)

# Remember Go env settings, so we can set them back
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
GOMOD=$(go env GO111MODULE)
GOCGO=$(go env CGO_ENABLED)

echo "[INF] LDFLAGS: ${LDFLAGS[*]}"

echo "[INF] setting build env"
 go env -w "CGO_ENABLED=0"

fmt() {
	echo "[INF] FMT"
	cd "${ROOTDIR}"
	go fmt .
	wait
	cd "${ROOTDIR}/bin"
	go fmt .
	cd "${ROOTDIR}"
	wait
}

build() {
	echo "[INF] make & install ${APPNAME}"
	fmt
	echo "[INF] Building"
	go build -o="${APPNAME}" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/bin/"
	wait
}

release() {
	echo "[INF] release ${APPNAME}"
	fmt
	echo "[INF] Building"
	go build -o="${APPNAME}-${GOOS}-${GOARCH}" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/bin/"
	wait
	env GOOS=linux GOARCH=arm64	go build -o="${APPNAME}-linux-arm64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/bin/"
	wait
	env GOOS=windows GOARCH=amd64 go build -o="${APPNAME}-windows-amd64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/bin/"
	wait
	env GOOS=darwin GOARCH=amd64	go build -o="${APPNAME}-darwin-amd64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/bin/"
	wait
}



# executes the function $1 $2 $3 $4 $5 are the arguments
$@

# Wait for the functions to finsih
wait

echo "[INF] setting Go env back to:"
echo "[INF] GOARCH=${GOARCH} GOOS=${GOOS} GO111MODULE=${GOMOD} GOCGO=${GOCGO}"
go env -w GOARCH=${GOARCH} GOOS=${GOOS} CGO_ENABLED=${GOCGO}
go env -w GO111MODULE=${GOMOD}
echo "[INF] all done"
