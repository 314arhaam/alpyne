# `alpyne` Deep Dive into Docker Images

## How to:

1. Clone this repo
2. Create `data/` in the root

### Decompose

1. In the root, run `./scripts/decompose.sh -v --im <IMAGE_NAME>`. You might need `sudo`
2. Check `data/`

### Build

1. In the root, run `./scripts/build.sh -v --im <DECOMPOSED_IMAGE_DIRECTORY>`. You might need `sudo`
2. You'll be in the container
