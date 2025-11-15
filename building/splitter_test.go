package building_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/RyanLambrecht/catwalk/building"
)

func totalOutput(s building.Splitter) float64 {
	_, ratios := s.Output()
	total := s.Available()

	for _, r := range ratios {
		total += r
	}
	return total
}

func expectPanic(t *testing.T, should bool, fn func()) {
	t.Helper()
	defer func() {
		r := recover()
		if should && r == nil {
			t.Fatalf("expected panic")
		}
		if !should && r != nil {
			t.Fatalf("unexpected panic: %v", r)
		}
	}()
	fn()
}

// should be overkill as only this many links should happen for splitting
// concreate from the global pool, yet still overkill
func TestSetOutput_LargeLog(t *testing.T) {
	NumberOfTests := 500

	for i := 1; i <= NumberOfTests; i += (NumberOfTests / 10) - 1 {
		t.Run(fmt.Sprintf("%d outputs", i), func(t *testing.T) {
			s := building.NewSplitter(nil, i)
			for j := range i {
				s.SetOutput(j, nil, s.Available()/2)
			}

			const eps = 1e-6
			if math.Abs(1.0-totalOutput(s)) > eps {
				t.Errorf("expected 1.0, got %f", totalOutput(s))
			}
		})

	}
}

func TestSetOutput_CorrectTotal(t *testing.T) {
	//l := &building.Link{}

	tests := []struct {
		name      string
		wantPanic bool
		outputs   int
		ratios    []float64
	}{
		{
			name:      "assign no output",
			wantPanic: false,
			outputs:   1,
			ratios:    []float64{},
		},
		{
			name:      "small, intermediate",
			wantPanic: false,
			outputs:   2,
			ratios:    []float64{0.5, 0},
		},
		{
			name:      "normal, uniform",
			wantPanic: false,
			outputs:   5,
			ratios:    []float64{0.2, 0.2, 0.2, 0.2, 0.2},
		},
		{
			name:      "normal, intermediate",
			wantPanic: false,
			outputs:   5,
			ratios:    []float64{0.2, 0, 0, 0, 0},
		},
		{
			name:      "over allocation",
			wantPanic: true,
			outputs:   3,
			ratios:    []float64{1.0 / 3, 2.0 / 3, 1.0 / 3}, // panics on third
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {
			expectPanic(t, tc.wantPanic, func() {
				s := building.NewSplitter(nil, tc.outputs)
				for i, ratio := range tc.ratios {
					s.SetOutput(i, nil, ratio)
				}

				const eps = 1e-6
				if !tc.wantPanic {
					if math.Abs(1.0-totalOutput(s)) > eps {
						t.Errorf("expected total 1, got %f", totalOutput(s))
					}
				}
			})
		})

	}

}
