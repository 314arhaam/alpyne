package tartools

import (
	// "os"
	"os/exec"
	"fmt"
	"local/alpyne/internal/iotools"
)

func UnTar(filename, dst string) error {
	if err := iotools.SafeMkdir(dst); err != nil {
		return fmt.Errorf("tarman.go: UnTar(...): Create directory stage %w", err)
	}
	cmd := exec.Command("tar", "-xvf", filename, "-C", dst)
	/*cmd.StdOut = os.StdOut
	cmd.stderr = os.stderr*/
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("tarman.go: UnTar(...): Command run stage %w", err)
	}
	return nil
}