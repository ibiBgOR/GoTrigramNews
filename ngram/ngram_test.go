package ngram

import "testing"

func TestNgram(t *testing.T) {
	cases := []struct {
		n    int
		str  string
		want []string
	}{
		{3, "Hello world", []string{"Hel", "ell", "llo", "lo ", "o w", " wo", "wor", "orl", "rld"}},
		{2, "Hello world", []string{"He", "el", "ll", "lo", "o ", " w", "wo", "or", "rl", "ld"}},
	}

	for _, c := range cases {
		err := false
		got := BuildNGram(c.str, c.n)
		if len(got) == len(c.want) {
			for index, _ := range got {
				if got[index] != c.want[index] {
					err = true
					//t.Errorf("item %q abd %q differ", got[index], c.want[index])
				}
			}
		} else {
			err = true
			t.Errorf("length differ: %q %q", string(len(got)), string(len(c.want)))
		}
		if err {
			t.Errorf("BuildNGram(%q, %q): \n got:  %q,\n want: %q", c.str, string(c.n), got, c.want)
		}
	}
}
