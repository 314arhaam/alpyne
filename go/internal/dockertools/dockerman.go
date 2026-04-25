package dockertools

import (
	"github.com/moby/moby/client"
	"context"
	"fmt"
	"os"
	"io"
)

func CreateClient() (*client.Client, error) {
	apiClient, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("dockerman.go: CreateClient(...): %w", err)
	}
	return apiClient, nil
}

func SaveImageToTar(cli *client.Client, ImageName, dst string) error {
	ctx := context.Background()
	imageName := []string{ImageName}
	// read image data
	ImageData, err := cli.ImageSave(ctx, imageName)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer ImageData.Close()
	// create file to save
	file, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	defer file.Close()
	// copy data to file
	if _, err := io.Copy(file, ImageData); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}