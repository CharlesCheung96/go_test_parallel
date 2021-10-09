package pkg1

import (
	"testdemo/demo"
	"testing"
)

func TestPkg1(t *testing.T) {
	t.Parallel()
	for i := 1; i < 10000; i++ {
		demo.AccessOuter("pkg1", "TestPkg1")
	}
}

func TestProvider(t *testing.T) {
	t.Parallel()
	demo.Provider()
}
