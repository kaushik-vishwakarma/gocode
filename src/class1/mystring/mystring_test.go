package mystring

import "testing"

func TestLength(t *testing.T) {
	var tests = []struct {
		s string
		l int
	} {
		 {"kaushik", 7 },
		 {"kavi", 4 },
		 {"", 0 },
	}

	for _, c  := range tests {
		got := Length(c.s)
		if got != c.l {
			t.Errorf("Length(%q) == %q, want %q", c.s, got, c.l)
		}
	}
}

func TestEven(t *testing.T) {
	var tests = []struct {
		s string
		e bool
	} {
		 {"kaushik", false },
		 {"kavi", true },
		 {"", true },
	}

	for _, c  := range tests {
		got, e := Iseven(c.s)
		if got != c.e && e == nil {
			t.Errorf("Iseven(%q) == %q, want %q", c.s, got, c.e)
		}
	}
}

func TestEven2(t *testing.T) {
	var tests = []struct {
		s string
		e bool
	} {
		 {"kaushik", false },
		 {"kavi", true },
		 {"", true },
	}

	for _, c  := range tests {
		got, e := Iseven2(c.s)
		if got != c.e && e == nil {
			t.Errorf("Iseven(%q) == %q, want %q", c.s, got, c.e)
		}
	}
}
