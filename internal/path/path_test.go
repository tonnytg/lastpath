package path_test

import (
	"lastpath/internal/path"
	"testing"
)

func TestNewPath(t *testing.T) {
	p := path.Path{}
	p.NewPath()
	if p.Id != 1 {
		t.Errorf("ID is not 1, it is: %d", p.Id)
	}
}
