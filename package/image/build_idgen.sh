#!/bin/sh

set -x

# Run this script: buildah unshare ./build.sh

echo "Build idgen image using buildah"

BUILDAH_HISTORY=true
NAME=docker.io/mikegolovanov/idgen
VERSION=1.0 

container=$(buildah from scratch)
echo "Create container $container"
#mnt=$(buildah mount $container)
#echo "Mount container filesystem to $mnt"
buildah add $container bin/users_idgen /idgen

buildah config --cmd "/idgen" $container
buildah config --port 8080 $container 

img=$(buildah commit $container)
#buildah umount $container


buildah tag $img $NAME:$VERSION
echo "Create image $img done"
 
buildah login -u mikegolovanov docker.io
buildah push $img 
buildah logout docker.io