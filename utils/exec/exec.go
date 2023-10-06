package executils

import (
	"os"
	"os/exec"
)

func RunWithStdIo(name string, args ...string) error {
	return CommandWithStdIo(name, args...).Run()
}

func CommandWithStdIo(name string, args ...string) *exec.Cmd {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd
}

func CheckAvailableCommands(cmds ...string) error {
	for _, c := range cmds {
		if _, err := exec.LookPath(c); err != nil {
			return err
		}
	}
	return nil
}
