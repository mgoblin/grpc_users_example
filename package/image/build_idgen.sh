#!/bin/sh

set -x

# Run this script: buildah unshare ./build.sh

echo "Build idgen image using buildah"

BUILDAH_HISTORY=true
NAME=docker.io/mikegolovanov/idgen
VERSION=1.0 

container=$(buildah from scratch)
echo "Create container $container"
buildah add $container bin/users_idgen /idgen

buildah config --cmd "/idgen" $container
buildah config --port 8080 $container 

img=$(buildah commit $container)

buildah tag $NAME:$VERSION
echo "Create image $img done"
 
buildah login -u mikegolovanov docker.io
buildah push $img $NAME:$VERSION 
buildah logout docker.io