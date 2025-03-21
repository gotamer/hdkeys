#!/usr/bin/env bash

set -e
clear

APPNAME='hdkeys'
PACKAGE="go.hansaray.pw/${APPNAME}"

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
GOMOD=$(go env GO111MODULE)
GOCGO=$(go env CGO_ENABLED)

echo "[INF] LDFLAGS: ${LDFLAGS[*]}\n"

echo "[INF] setting build env"
go env -w "CGO_ENABLED=0"

# Just edit Title, Description, and Tags
metadata() {
    echo '# Do not edit this file here, see make.sh' > metadata.toml
    echo "Name = '${APPNAME}'" >> metadata.toml 
    echo "Version = '${VERSION}'" >> metadata.toml 
    echo "Package = '${PACKAGE}'" >> metadata.toml 
    echo "Repo = 'https://github.com/gotamer/${APPNAME}'" >> metadata.toml 
    echo "Branch = 'master'" >> metadata.toml 
    echo "Homepage = 'http://${PACKAGE}'" >> metadata.toml 
    echo 'Title = "Hierarchical Deterministic Keys"' >> metadata.toml 
    echo 'Description = "A library to create Bitcoin and Nostr keys from the same mnemonic seeds, and a command-line tool that uses this library"' >> metadata.toml 
    echo "Documentation = 'https://pkg.go.dev/${PACKAGE}'" >> metadata.toml 
    echo 'Tags = ["nostr", "nips", "NIP05", "NIP06", "NIP19", "bitcoin", "BIP32", "BIP39", "BIP43", "BIP44", "BIP84", "BIP86", "BIP173", "SLIP44"]' >> metadata.toml 
    echo 'License = "MIT"' >> metadata.toml 
}

fmt() {
	echo "[INF] FMT"
	go fmt "${ROOTDIR}/cmd"
	wait
	go fmt "${ROOTDIR}/lib"
	wait
}

build() {
	echo "[INF] make & install ${APPNAME}"
	fmt
	echo "[INF] Building"
	go build -o="bin/${APPNAME}" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/cmd/"
	wait
}

release() {
	# go tool dist list | grep linux
	echo "[INF] release, fmt, metadata, build ${APPNAME}"
	fmt
	wait
	env GOOS=linux GOARCH=amd64 go build -o="bin/${APPNAME}-linux-amd64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/cmd/"
	wait
	env GOOS=linux GOARCH=arm64	go build -o="bin/${APPNAME}-linux-arm64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/cmd/"
	wait
	env GOOS=windows GOARCH=amd64 go build -o="bin/${APPNAME}-windows-amd64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/cmd/"
	wait
	env GOOS=darwin GOARCH=amd64 go build -o="bin/${APPNAME}-darwin-amd64" -ldflags="${LDFLAGS[*]}" "${ROOTDIR}/cmd/"
	wait
    metadata
	wait
}

# executes the function $1 $2 $3 $4 $5 are the arguments
$@

# Wait for the functions to finish
wait

echo "[INF] setting Go env back to:"
echo "[INF] GOARCH=$(go env GOHOSTARCH) GOOS=$(go env GOHOSTOS) GO111MODULE=${GOMOD} GOCGO=${GOCGO}"
go env -w GOARCH=$(go env GOHOSTARCH) 
go env -w GOOS=$(go env GOHOSTOS) 
go env -w CGO_ENABLED=${GOCGO}
go env -w GO111MODULE=${GOMOD}
echo "[INF] all done"
