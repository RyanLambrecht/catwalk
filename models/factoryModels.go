package models

// limited to one item type (we pray not one recipe)
type Link struct {
	To   *Node
	From *Node
}

type Splitter struct {
	input        *Link
	output       []*Link
	outputRatios []float32 //this + available = 1.0
	available    float32   //1.0 is all output avalabile
}

type Merger struct {
	input  []*Link
	output *Link
}

// this is a production building where the production multiplier is how
// many effetive buildings you need
// i.e. PM = 0.5 means its one building underlocked
//
//	PM = 2.0 means its either two buildings at 100% or one overclocked
type Node struct {
	buildingType         string
	ProductionMultiplier float32
	Recipe               string
	input                []*Link
	output               []*Link
}
