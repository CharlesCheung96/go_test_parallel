package ch1

import (
	"testdemo/demo1/chdemo"
	"testing"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	chdemo.Provider()
}
