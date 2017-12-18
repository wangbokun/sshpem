package main

import (
	"easyWork/ops/skm"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/fatih/color"
)

func Execute(workDir, script string, args ...string) bool {
	cmd := exec.Command(script, args...)

	if workDir != "" {
		cmd.Dir = workDir
	}

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		color.Red("%s%s", CrossSymbol, err.Error())
		return false
	}

	return true
}

func main() {
	args := []string{}
	args = append(args, "-t")
	args = append(args, "rsa")

	args = append(args, "-P")
	args = append(args, "")

	args = append(args, "-C")
	args = append(args, "wangbokun")

	args = append(args, "-f")
	args = append(args, filepath.Join("/tmp/", "", "wangbokun"))

	Execute("", "ssh-keygen", args...)
	color.Green("%sSSH key [%s] created!", skm.CheckSymbol, "alias")

}
