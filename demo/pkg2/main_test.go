package pkg1

import (
	"testdemo/demo/demo"
	"testing"
)

func TestPkg2(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		demo.AccessOuter("pkg2", "TestPkg2")
	}
}