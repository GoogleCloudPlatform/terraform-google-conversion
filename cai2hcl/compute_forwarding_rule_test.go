package cai2hcl

import (
	"testing"
)

func TestConvertComputeForwardingRule(t *testing.T) {
	cases := []struct {
		name string
	}{
		{name: "full_compute_forwarding_rule"},
	}
	for i := range cases {
		c := cases[i]

		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			err := assertTestData(c.name)
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
