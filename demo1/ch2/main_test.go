package ch2

import (
	"testdemo/demo1/chdemo"
	"testing"
)

func TestConsumer(t *testing.T) {
	t.Parallel()
	chdemo.Consumer()
}