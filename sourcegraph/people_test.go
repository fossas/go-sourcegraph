package sourcegraph

import "testing"

func TestPersonSpec(t *testing.T) {
	tests := []struct {
		str  string
		spec PersonSpec
	}{
		{"a", PersonSpec{Login: "a"}},
		{"a@a.com", PersonSpec{Email: "a@a.com"}},
		{"$1", PersonSpec{UID: 1}},
	}

	for _, test := range tests {
		spec, err := ParsePersonSpec(test.str)
		if err != nil {
			t.Errorf("%q: ParsePersonSpec failed: %s", test.str, err)
			continue
		}
		if spec != test.spec {
			t.Errorf("%q: got spec %+v, want %+v", test.str, spec, test.spec)
			continue
		}

		str := test.spec.PathComponent()
		if str != test.str {
			t.Errorf("%+v: got str %q, want %q", test.spec, str, test.str)
			continue
		}
	}
}
