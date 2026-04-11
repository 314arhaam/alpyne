#!/bin/bash
source scripts/.colors
set -euo pipefail

VERBOSE=0
while [[ $# -gt 0 ]]; do
  key="$1"
  case $key in
    -i|--image)
      CUSTOM_IMAGE="$2"
      echo "  - Building Image: $CUSTOM_IMAGE"
      shift
      shift
      ;;
    -v|--verbose)
      VERBOSE=1
      shift
      ;;
  esac
done

echo -e ""
mkdir data/$CUSTOM_IMAGE/STACK
for f in $(ls data/$CUSTOM_IMAGE/ -U | grep -v metadata | sort -V | xargs); do
echo -e "${GREEN} $f ${RESET}"
    cp -r data/$CUSTOM_IMAGE/$f/* data/$CUSTOM_IMAGE/STACK/ -f
done
echo -e "[*] STEP 07 - BUILD CONTAINER & RUN"
if [ "$VERBOSE" == "1" ]; then
    echo -e "${YELLOW}   -> Copy Dockerfile and .dockerignore to data/$CUSTOM_IMAGE/STACK/${RESET}"
fi
cp Dockerfile .dockerignore data/$CUSTOM_IMAGE/STACK/
if [ "$VERBOSE" == "1" ]; then
    echo -e "${YELLOW}   -> Building Image...${RESET}"
fi
docker build data/$CUSTOM_IMAGE/STACK -t "$CUSTOM_IMAGE:latest"
if [ "$VERBOSE" == "1" ]; then
    echo -e "${YELLOW}   -> Running container${RESET}"
fi
docker run -it --rm "$CUSTOM_IMAGE:latest"