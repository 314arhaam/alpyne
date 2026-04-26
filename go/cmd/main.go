package main

import (
	dt "local/alpyne/internal/dockertools"
	tt "local/alpyne/internal/tartools"
	"local/alpyne/internal/iotools"
	"fmt"
	"strconv"
	"os"
	"strings"
)

func main() {
	imageName := "docker.iranserver.com/alpine:latest"
	dirName := strings.Map(func(r rune) rune {
		if r == '.' || r == '/' || r == ':' {
			return '_'
		}
		return r
	}, imageName)
	cli, err := dt.CreateClient()
	if err != nil {
		fmt.Println()
		return
	}
	if err := dt.SaveImageToTar(cli, imageName, "alp.tar"); err != nil {
		fmt.Println(err)
		return
	}
	if err := tt.UnTar("alp.tar", dirName); err != nil {
		fmt.Println(err)
		return
	}
	var man []dt.Manifest
	if err := dt.ReadManifest(&man, dirName + "/manifest.json"); err != nil {
		fmt.Printf("%w", err)
		return
	}
	fmt.Println(man)
	fmt.Println(len(man[0].Layers))
	for i, f := range man[0].Layers {
		folderName := dirName + "/LAYER_" + strconv.Itoa(i)
		layerFile := dirName + "/" + f
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
	iotools.SafeMkdir(dirName + "/metadata")
	a, _ := os.ReadDir(dirName + "/blobs/sha256")
	for _, i := range a {
		fmt.Println(i.Name())
		os.Rename(dirName + "/blobs/sha256/" + i.Name(), dirName + "/metadata/" + i.Name())
	}
	os.RemoveAll(dirName + "/blobs")
}