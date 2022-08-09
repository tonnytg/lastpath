package path_test

import (
	"lstp/internal/path"
	"testing"
)

func TestNewPath(t *testing.T) {
	p := path.Program{}
	p.NewPath()
	if p.Id != 1 {
		t.Errorf("ID is not 1, it is: %d", p.Id)
	}
}
