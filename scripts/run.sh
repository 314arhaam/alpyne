#!/bin/bash

GREEN="\033[32m"
BLUE="\033[34m"
YELLOW="\033[33m"
RED="\033[31m"
RESET="\033[0m"

IMG_NAME="alpine"
TAG="latest"
MIRROR="docker.iranserver.com"
CUSTOM_IMAGE="custom-alpine"

# step 1 - if alpine image exists, passes. if not, pulls from image registry
echo -e "[*] STEP 01 - PULLING $IMG_NAME FROM $MIRROR"
if [ "$(docker images $MIRROR/$IMG_NAME:$TAG -q)" == "" ]; then
    echo -e "${RED}  -> Image not exists; Pulling...${RESET}"
    docker pull $MIRROR/$IMG_NAME:$TAG
else
    echo -e "${GREEN}  -> Image already exists. Using it.${RESET}"
fi

# step 2 - save image as a tar file to data/alpine file
echo -e "[*] STEP 02 - SAVING $MIRROR/$IMG_NAME TO aplyne/data/apline"
rm -rf data
mkdir data
docker image save $MIRROR/$IMG_NAME:$TAG -o data/$IMG_NAME

# step 3 - decopress tar files.
echo -e "[*] STEP 03 - UNTAR IMAGE FILES"
tar -xf data/$IMG_NAME -C data/
if [ "$1" == "-v" ]; then
    echo -e "${YELLOW}   -> View manifest.json:${RESET}"
    cat data/manifest.json | jq
    echo -e "${YELLOW}\n Tree of directory:${RESET}"
    tree data/
fi

echo -e "[*] STEP 04 - UNTAR MAIN LAYERS"
mkdir data/alpyne
LAYER_TAR_ADDR=$(cat data/manifest.json | jq .[].Layers[] -r)
LAYER_TAR_FNAME=$(echo $LAYER_TAR_ADDR | awk -F/ '{print $3}')
echo "  -> Layer filename: $LAYER_TAR_FNAME"
mv data/$LAYER_TAR_ADDR data/alpyne/
tar -xf data/alpyne/$LAYER_TAR_FNAME -C data/alpyne

echo -e "[*] STEP 05 - MODIFICATION"
if [ "$1" == "-v" ]; then
    echo -e "${YELLOW}${YELLOW}   -> Original os-release file:${RESET}"
    cat data/alpyne/etc/os-release
    echo -e "${YELLOW}${YELLOW}   -> Modified os-release file:${RESET}"
    cat os-release
fi
cp os-release data/alpyne/etc/os-release

echo -e "[*] STEP 06 - CLEANUP"
rm -rf data/alpyne/$LAYER_TAR_FNAME
rm -rf data/alpine  data/blobs  data/index.json  data/manifest.json  data/oci-layout
mv data/alpyne/* data/
rm -rf data/alpyne

echo -e "[*] STEP 07 - BUILD CONTAINER & RUN"
if [ "$1" == "-v" ]; then
    echo -e "${YELLOW}   -> Copy Dockerfile and .dockerignore to data/${RESET}"
fi
cp Dockerfile .dockerignore data/
if [ "$1" == "-v" ]; then
    echo -e "${YELLOW}   -> Building Image...${RESET}"
fi
docker build data/ -t "$CUSTOM_IMAGE:latest"
if [ "$1" == "-v" ]; then
    echo -e "${YELLOW}   -> Running container${RESET}"
fi
docker run -it --rm "$CUSTOM_IMAGE:latest"