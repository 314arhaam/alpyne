package tartools

import (
	// "os"
	"os/exec"
	"fmt"
	"local/alpyne/internal/iotools"
)

func UnTar(filename, dst string) error {
	if err := iotools.SafeMkdir(dst); err != nil {
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