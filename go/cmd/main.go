package main

import (
	dt "local/alpyne/internal/dockertools"
	tt "local/alpyne/internal/tartools"
	ct "local/alpyne/internal/clitools"
	"local/alpyne/internal/iotools"
	"fmt"
	"strconv"
	"os"
	"strings"
	"path"
)

func main() {
	cliargs := ct.CLIArgs{}
	cliargs.Init()
	imageName := (*cliargs.ImageName)
	dirName := "IMG_" + strings.Map(func(r rune) rune {
		if r == '.' || r == '/' || r == ':' {
			return '_'
		}
		return r
	}, imageName)
	imageTar := dirName + ".tar"
	// Print some stuff
	fmt.Println("[*] Processing Image:		", imageName)
	cli, err := dt.CreateClient()
	if err != nil {
		fmt.Println("main.go:", err)
		return
	}
	if err := dt.SaveImageToTar(cli, imageName, imageTar); err != nil {
		fmt.Println("main.go:", err)
		return
	}
	fmt.Println("	-> Tar file:		", imageTar)
	if !(*cliargs.Keep){
		defer os.Remove(imageTar)
	}
	if err := tt.UnTar(imageTar, dirName); err != nil {
		fmt.Println("main.go:", err)
		return
	}
	fmt.Println("	-> Contents dir:	", dirName)
	var man []dt.Manifest
	fmt.Println("[*] Processing Layers")
	if err := dt.ReadManifest(&man, path.Join(dirName,"manifest.json")); err != nil {
		fmt.Println("main.go:", err)
		return
	}
	// fmt.Println(man)
	// fmt.Println(len(man[0].Layers))
	for i, layerSHA256Name := range man[0].Layers {
		folderName := path.Join(dirName, "LAYER_" + strconv.Itoa(i+1))
		layerFile := path.Join(dirName, layerSHA256Name)
		// fmt.Println(folderName, layerSHA256Name)
		fmt.Printf("	-> Layer %3d of %3d @ %s\n", i+1, len(man[0].Layers), folderName)
		if err := tt.UnTar(layerFile, folderName); err != nil {
			fmt.Println("main.go: error in deep un-tar", err)
			return
		}
		if err := os.Remove(layerFile); err != nil {
			fmt.Println("main.go: error in deep un-tar - Removal", err)
			return
		}
	}
	fmt.Println("[*] Processing Metadata")
	iotools.SafeMkdir(path.Join(dirName, "metadata"))
	files, _ := os.ReadDir(path.Join(dirName, "blobs/sha256"))
	for _, i := range files {
		fmt.Printf("	-> Moving blobs/sha256/%s to metadata/\n", i.Name())
		os.Rename(
			path.Join(dirName, "blobs/sha256", i.Name()), 
			path.Join(dirName, "metadata", i.Name()),
		)
	}
	fmt.Println("[*] Clean Up")
	os.RemoveAll(path.Join(dirName, "blobs"))
	//fmt.Println("Decomposed finished, data available in: ", dirName)
}