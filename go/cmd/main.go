package main

import (
	dt "local/alpyne/internal/dockertools"
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
}