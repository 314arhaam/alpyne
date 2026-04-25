package dockertools

import (
	dt "github.com/moby/moby/client"
)

func CreateClient() (*client.Client, error) {
	apiClient, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("dockerman.go: CreateClient(...): %w", err)
	}
	return apiClient, nil
}

func SaveImageToTar() {
	// read image data
	// create file to save
	// copy data to file
}