#!/bin/bash
source scripts/.colors

while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
    -i|--image)
      IMG_FULL_NAME="$2"
      echo "  - Inspecting Image: $IMG_FULL_NAME"
      shift
      shift
      ;;
    -v|--verbose)
      VERBOSE=1
      shift
      ;;
  esac
done

IMG_NAME=$(echo $IMG_FULL_NAME | sed s/"\/"/"__IMG__"/g | sed s/":"/"__VER__"/g | sed s/"\."/_/g)
CUSTOM_IMAGE="custom-$IMG_NAME"

# step 1 - if alpine image exists, passes. if not, pulls from image registry
echo -e "[*] STEP 01 - PULLING $IMG_FULL_NAME"
if [ "$(docker images $IMG_FULL_NAME -q)" == "" ]; then
    echo -e "${RED}  -> Image not exists; Pulling...${RESET}"
    docker pull $IMG_FULL_NAME
else
    echo -e "${GREEN}  -> Image already exists. Using it.${RESET}"
fi

# step 2 - save image as a tar file to data/alpine file
echo -e "[*] STEP 02 - SAVING $IMG_NAME TO aplyne/data/$IMG_NAME"
rm -rf data/$IMG_NAME data/$CUSTOM_IMAGE
mkdir data/$IMG_NAME data/$CUSTOM_IMAGE
docker image save $IMG_FULL_NAME -o data/$IMG_NAME.dockerimage

# step 3 - decopress tar files.
echo -e "[*] STEP 03 - UNTAR IMAGE FILES"
tar -xf data/$IMG_NAME.dockerimage -C data/$IMG_NAME
if [ "$VERBOSE" == "1" ]; then
    echo -e "${YELLOW}   -> View manifest.json:${RESET}"
    cat data/$IMG_NAME/manifest.json | jq
    echo -e "${YELLOW}\n Tree of directory:${RESET}"
    tree data/$IMG_NAME
fi

# step 4 - export layers and untar data
echo -e "[*] STEP 04 - UNTAR MAIN LAYERS"
layer_index=0
for LAYER_TAR_ADDR in $(cat data/$IMG_NAME/manifest.json | jq .[].Layers[] -r | xargs); do
    LAYER_TAR_FNAME=$(echo $LAYER_TAR_ADDR | awk -F/ '{print $3}');
    echo "  -> Layer filename: $LAYER_TAR_FNAME";
    mv data/$IMG_NAME/$LAYER_TAR_ADDR data/$CUSTOM_IMAGE/;
    mkdir data/$CUSTOM_IMAGE/LAYER_$layer_index;
    tar -xf data/$CUSTOM_IMAGE/$LAYER_TAR_FNAME -C data/$CUSTOM_IMAGE/LAYER_$layer_index;
    rm -rf data/$CUSTOM_IMAGE/$LAYER_TAR_FNAME
    layer_index=$((layer_index+1))
done
mkdir data/$CUSTOM_IMAGE/metadata
mv data/$IMG_NAME/blobs/sha256/* data/$CUSTOM_IMAGE/metadata
mv data/$IMG_NAME/index.json  data/$IMG_NAME/manifest.json  data/$IMG_NAME/oci-layout data/$CUSTOM_IMAGE/metadata/
rm -rf data/$IMG_NAME data/$IMG_NAME.*
tree data/$CUSTOM_IMAGE -L 1
tree data/$CUSTOM_IMAGE/metadata