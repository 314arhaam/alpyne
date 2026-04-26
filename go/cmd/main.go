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
	if err := tt.UnTar("alp.tar", "alpdata"); err != nil {
		fmt.Println()
		return
	}
}