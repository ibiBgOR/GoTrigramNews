package data

import "testing"

func TestCreateDatabase(t *testing.T) {
	createDatabase("root", "")
}

func TestGetIds(t *testing.T) {
	cases := []struct {
		in   string
		want []int
	}{
		{"Hello world", []int{}},
	}

	for _, c := range cases {
		err := false
		got := getIds(c.in)
		if len(got) == len(c.want) {
			for index, _ := range got {
				if got[index] != c.want[index] {
					err = true
				}
			}
		} else {
			err = true
		}
		if err {
			t.Errorf("getIds(%q): \n got:  %q,\n want: %q", c.in, got, c.want)
		}
	}
}
