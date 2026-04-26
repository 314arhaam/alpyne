package iotoools

import (
	"os"
	"fmt"
)

func SafeMkdir(dst string) error {
	switch err := os.Mkdir(dst, 0o755); err != nil {
	case os.IsExist(err):
		fmt.Println("File exists, re-write")
		return nil
	default:
		return fmt.Errorf("%w", err)
	}
}