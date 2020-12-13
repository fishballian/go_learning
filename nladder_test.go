package go_learning

import "testing"

func TestNLadder(t *testing.T) {
	cases := []struct {
		Input int
		Output int
	}{
		{1,1},
		{2,2},
		{3, 3},
	}
	for _, c := range cases {
		out := NLadder(c.Input)
		if out != c.Output {
			t.Fatalf("NLadder(%d), want %d but got %d", c.Input, c.Output, out)
		}
	}
	t.Log(NLadder(100))
}
