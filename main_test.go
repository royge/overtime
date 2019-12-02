package main

import (
	"fmt"
	"testing"
)

func TestCalculte(t *testing.T) {
	workHours := "9h30m"

	tt := []struct {
		start    string
		end      string
		wh       string
		overtime string
	}{
		{
			"10:51",
			"21:13",
			workHours,
			"52m0s",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("%s - %s", tc.start, tc.end), func(t *testing.T) {
			res, err := calculate(tc.start, tc.end, tc.wh)
			if err != nil {
				t.Errorf("expected no error, got %v", err)
			}

			if res.String() != tc.overtime {
				t.Errorf("expected %s, got %s", tc.overtime, res)
			}
		})
	}
}
