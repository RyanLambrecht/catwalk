package building

import "fmt"

type splitter struct {
	input        *Link
	output       []*Link
	outputRatios []float64 //this + available = 1.0
	available    float64   //1.0 is all output avalabile
}

// ideally, this will be used to optomize/ unuion anything that uses the
// same recipe
type Splitter interface {
	Available() float64
	Output() ([]*Link, []float64)
	SetOutput(index int, link *Link, ratio float64)
	Input() *Link
}

// should only be used when the input and output are known
func NewSplitter(input *Link, outputCount int) Splitter {
	return &splitter{
		available:    1,
		output:       make([]*Link, outputCount),
		outputRatios: make([]float64, outputCount),
		input:        input,
	}
}

func (s *splitter) Available() float64 {
	return s.available
}

// returns outputLinks, outputRatios
func (s *splitter) Output() ([]*Link, []float64) {
	return s.output, s.outputRatios
}

func (s *splitter) SetOutput(index int, link *Link, ratio float64) {
	delta := ratio - s.outputRatios[index]
	if delta > s.available {
		panic(fmt.Sprintf("internal error: ratio %.2f exceeds available %.2f with difference %.2f in splitter for recipe %s",
			ratio, s.available, delta, link.From.Recipe))
	}

	s.output[index] = link
	s.outputRatios[index] = ratio
	s.available -= delta
}

func (s *splitter) Input() *Link {
	return s.input
}
