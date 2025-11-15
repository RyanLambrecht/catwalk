package building

// limited to one item type (we pray not one recipe)
type Link struct {
	To   *Node
	From *Node
}

// this is a production building where the production multiplier is how
// many effetive buildings you need
// i.e. PM = 0.5 means its one building underlocked
//
//	PM = 2.0 means its either two buildings at 100% or one overclocked
type Node struct {
	buildingType         string
	ProductionMultiplier float64
	//maybe this can be used as the id for the building?
	Recipe string
	input  []*Link
	output []*Link
}
