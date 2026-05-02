package dockertools

import (
	"os"
	"fmt"
	"encoding/json"
)

type Manifest struct {
	Config		string		`json:"Config"`
	RepoTags	[]string	`json:"RepoTags"`
	Layers		[]string	`json:"Layers"`
}

func ReadManifest(m *[]Manifest, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("manifestor.go: ReadManifest(...): Create file stage: %w", err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(m); err != nil {
		return fmt.Errorf("manifestor.go: ReadManifest(...): Decoding stage: %w", err)
	}
	return nil
}