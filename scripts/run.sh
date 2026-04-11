#!/bin/bash

IMG_NAME="alpine"
TAG="latest"
MIRROR="docker.iranserver.com"
CUSTOM_IMAGE="custom-alpine"

echo -e "[*] STEP 01 - PULLING $IMG_NAME FROM $MIRROR"
if [ "$(docker images $MIRROR/$IMG_NAME:$TAG -q)" == "" ]; then
    echo "  -> Image not exists; Pulling..."
    docker pull $MIRROR/$IMG_NAME:$TAG
else
    echo "  -> Image already exists. Using it."
fi


echo -e "[*] STEP 02 - SAVING $MIRROR/$IMG_NAME TO aplyne/data/apline"
rm -rf data
mkdir data
docker image save $MIRROR/$IMG_NAME:$TAG -o data/$IMG_NAME

echo -e "[*] STEP 03 - UNTAR IMAGE FILES"
tar -xf data/$IMG_NAME -C data/
# cat data/manifest.json | jq

echo -e "[*] STEP 04 - UNTAR MAIN LAYERS"
mkdir data/alpyne
LAYER_TAR_ADDR=$(cat data/manifest.json | jq .[].Layers[] -r)
LAYER_TAR_FNAME=$(echo $LAYER_TAR_ADDR | awk -F/ '{print $3}')
echo "  -> Layer filename: $LAYER_TAR_FNAME"
mv data/$LAYER_TAR_ADDR data/alpyne/
tar -xf data/alpyne/$LAYER_TAR_FNAME -C data/alpyne

echo -e "[*] STEP 05 - MODIFICATION"
# cat data/alpyne/etc/os-release
cp os-release data/alpyne/etc/os-release

echo -e "[*] STEP 06 - CLEANUP"
rm -rf data/alpyne/$LAYER_TAR_FNAME
rm -rf data/alpine  data/blobs  data/index.json  data/manifest.json  data/oci-layout
mv data/alpyne/* data/
rm -rf data/alpyne

echo -e "[*] STEP 07 - BUILD CONTAINER & RUN"
cp Dockerfile .dockerignore data/
docker build data/ -t "$CUSTOM_IMAGE:latest"
docker run -it --rm "$CUSTOM_IMAGE:latest"