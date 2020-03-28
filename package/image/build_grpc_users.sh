#!/bin/sh

# Run this script: buildah unshare ./build.sh

echo "Build grpc users server image using buildah"

BUILDAH_HISTORY=true
REGISTRY=docker.io
NAME=$REGISTRY/mikegolovanov/grpc-users-server
VERSION=1.0
SRC_BIN_PATH=bin/users_server
DEST_BIN_PATH=/server
TCP_PORT=7777

container=$(buildah from scratch)
echo "Create container $container"
buildah add $container $SRC_BIN_PATH $DEST_BIN_PATH

buildah config --cmd $DEST_BIN_PATH $container
buildah config --port $TCP_PORT $container 

img=$(buildah commit $container)

buildah tag $img $NAME:$VERSION
echo "Create image $img done"
 
buildah login -u mikegolovanov $REGISTRY
buildah push $img $NAME:$VERSION 
buildah logout $REGISTRY

buildah rm $container