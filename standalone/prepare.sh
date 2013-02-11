#!/bin/bash

GO_VERSION=go1.0.3.linux-386.tar.gz

cd ~
curl -O https://go.googlecode.com/files/$GO_VERSION
tar xzf $GO_VERSION
rm $GO_VERSION

export GOROOT=~/go
export GOPATH=~/m-lab.pipeline/standalone

PATH=~/go/bin:$PATH

mkdir -p package/init
go build pipeline

cp pipeline package
cp m-lab.pipeline/standalone/start.sh package/init
cp m-lab.pipeline/standalone/stop.sh package/init

tar -C package -cf mlab_pipeline.tar .

