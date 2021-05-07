package go_learning

import "testing"

func TestHello(t *testing.T) {
	t.Log(Hello())
}

func TestMod(t *testing.T) {
	t.Log((0 - 1 + 24*60) % (24 * 60))
}
