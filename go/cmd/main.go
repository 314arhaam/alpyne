package main

import (
	dt "local/alpyne/internal/dockertools"
	tt "local/alpyne/internal/tartools"
	"fmt"
	"strconv"
	"os"
)

func main() {
	cli, err := dt.CreateClient()
	if err != nil {
		fmt.Println()
		return
	}
	if err := dt.SaveImageToTar(cli, "docker.iranserver.com/alpine:latest", "alp.tar"); err != nil {
		fmt.Println(err)
		return
	}
	if err := tt.UnTar("alp.tar", "image_data"); err != nil {
		fmt.Println(err)
		return
	}
	var man []dt.Manifest
	if err := dt.ReadManifest(&man, "image_data/manifest.json"); err != nil {
		fmt.Printf("%w", err)
		return
	}
	fmt.Println(man)
	fmt.Println(len(man[0].Layers))
	for i, f := range man[0].Layers {
		folderName := "image_data/LAYER_" + strconv.Itoa(i)
		layerFile := "image_data/" + f
		fmt.Println(folderName, f)
		if err := tt.UnTar(layerFile, folderName); err != nil {
			fmt.Println("Error in deep un-tar", err)
			return
		}
		if err := os.Remove(layerFile); err != nil {
			fmt.Println("Error in deep un-tar - Removal", err)
			return
		}
	}
}