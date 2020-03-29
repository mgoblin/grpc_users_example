#!/bin/sh

# Run this script: buildah unshare ./build.sh

echo "Build idgen image using buildah"

BUILDAH_HISTORY=true
REGISTRY=docker.io
NAME=$REGISTRY/mikegolovanov/idgen
VERSION=1.0 

TCP_PORT=8080

container=$(buildah from scratch)
echo "Create container $container"
buildah add $container bin/users_idgen /idgen

buildah config --cmd /idgen $container
buildah config --port $TCP_PORT $container 

img=$(buildah commit $container idgen)

buildah tag $img $NAME:$VERSION
echo "Create image $img done"
 
buildah login -u mikegolovanov $REGISTRY
buildah push $img $NAME:$VERSION 
buildah logout $REGISTRY

buildah rm $container