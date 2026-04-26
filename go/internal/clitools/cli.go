package clitools

import (
	"flag"
)

type CLIArgs struct {
	ImageName	*string
	Keep		*bool
}

func (c *CLIArgs) Init() {
	c.ImageName = flag.String("i", "", "Name of the image to decompose")
	c.Keep = flag.Bool("k", false, "Keep image tar file")
	flag.Parse()
	if (*c.ImageName) == "" {
		panic("No image name provided. Use flag `-i` to pass image name.")
	}
}