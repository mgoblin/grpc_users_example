#!/bin/sh

# Run this script: buildah unshare ./build.sh

echo "Build users web client image using buildah"

BUILDAH_HISTORY=true
REGISTRY=docker.io
NAME=$REGISTRY/mikegolovanov/web-users-server
VERSION=1.0
SRC_BIN_PATH=bin/users_web
DEST_BIN_PATH=/server
TCP_PORT=8080

container=$(buildah from scratch)
echo "Create container $container"
buildah add $container $SRC_BIN_PATH $DEST_BIN_PATH
buildah add $container web/template/index.html web/template/index.html

buildah config --cmd $DEST_BIN_PATH $container
buildah config --port $TCP_PORT $container 

img=$(buildah commit $container users_webclient)

buildah tag $img $NAME:$VERSION
echo "Create image $img done"
 
buildah login -u mikegolovanov $REGISTRY
buildah push $img $NAME:$VERSION 
buildah logout $REGISTRY

buildah rm $container