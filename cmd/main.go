package main

import (
	"path/filepath"

	"github.com/fatih/color"
	"github.com/wangbokun/gowork/exec"
	"github.com/wangbokun/sshpem/models"
)

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

	exec.Execute("", "ssh-keygen", args...)
	color.Green("%sSSH key [%s] created!", models.CheckSymbol, "alias")

}
