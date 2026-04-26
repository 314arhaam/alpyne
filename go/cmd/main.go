package main

import (
	dt "local/alpyne/internal/dockertools"
	tt "local/alpyne/internal/tartools"
	"fmt"
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
		fmt.Println()
		return
	}
	var man []dt.Manifest
	if err := dt.ReadManifest(&man, "image_data/manifest.json"); err != nil {
		fmt.Printf("%w", err)
		return
	}
	fmt.Println(man)
	fmt.Println(len(man[0].Layers))
}