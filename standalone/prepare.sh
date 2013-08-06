#!/bin/bash

set -x 
set -e

if [ -z "$SOURCE_DIR" ] ; then
    echo "Expected SOURCE_DIR in environment"
    exit 1
fi
if [ -z "$BUILD_DIR" ] ; then
    echo "Expected BUILD_DIR in environment"
    exit 1
fi

if test -d $BUILD_DIR ; then
    rm -rf $BUILD_DIR/*
fi

# install dependencies such as development tools
yum groupinstall -y 'Development tools'

# NOTE: copy any files needed by the installed package
cp -r $SOURCE_DIR/init           $BUILD_DIR/

cat <<\EOF > $BUILD_DIR/conf/config.sh
RSYNCDIR_FATHOM=fathom
EOF
