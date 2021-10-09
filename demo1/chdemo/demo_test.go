package chdemo

import "testing"

func TestConsumer(t *testing.T) {
	t.Parallel()
	Consumer()
}

func TestProvider(t *testing.T) {
	t.Parallel()
	Provider()
}
