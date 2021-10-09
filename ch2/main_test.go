package ch2

import (
	"testdemo/chdemo"
	"testing"
)

func TestConsumer(t *testing.T) {
	t.Parallel()
	chdemo.Consumer()
}