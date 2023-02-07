package main_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestCmd(t *testing.T) {
	cmd := exec.Command("go", "build", "-o", "lastpath", ".")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
	defer os.Remove("lastpath")

	cmd = exec.Command("./lastpath", "save")
	if err := cmd.Run(); err != nil {
		t.Fatal(err)
	}
}
