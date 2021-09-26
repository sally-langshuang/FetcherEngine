package osencap

import (
	"fmt"
	"os/exec"
)
func F(args ...string) (string, error) {
	return Exec("echo", "123")
}

func Exec(cmd string, args ...string) (string, error) {
	cmdpath, err := exec.LookPath(cmd)
	if err != nil {
		fmt.Errorf("exec.LookPath err: %v, cmd: %s", err, cmd)
		return "", fmt.Errorf("infra.ErrExecLookPathFailed")
	}

	var output []byte
	output, err = exec.Command(cmdpath, args...).CombinedOutput()
	if err != nil {
		fmt.Errorf("exec.Command.CombinedOutput err: %v, cmd: %s", err, cmd)
		return "", fmt.Errorf("infra.ErrExecCombinedOutputFailed")
	}
	fmt.Println("CMD[", cmdpath, "]ARGS[", args, "]OUT[", string(output), "]")
	return string(output), nil
}