package cmd_test

import (
	"lstp/internal/cmd"
	"testing"
)

func TestCmd(t *testing.T) {

	err := cmd.CheckFlag()
	if err != nil {
		t.Error(err)
	}
}
