package ch1

import (
	"testdemo/chdemo"
	"testing"
)

func TestProvider(t *testing.T) {
	t.Parallel()
	chdemo.Provider()
}
