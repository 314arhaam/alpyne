package tartools

import (
	"os"
	"os/exec"
	"fmt"
)

func UnTar(filename, dst string) error {
	switch err := os.Mkdir(dst, 0o755); err != nil {
	case os.IsExist(err):
		fmt.Println("File exists, re-write")
	default:
		return fmt.Errorf("%w", err)
	}
	cmd := exec.Command("tar", "-xvf", filename, "-C", dst)
	/*cmd.StdOut = os.StdOut
	cmd.stderr = os.stderr*/
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}