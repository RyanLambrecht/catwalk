package models

func (m *Node) getBuildingType() string {
	return m.buildingType
}

func (s *Splitter) getAvailable() float32 {
	return s.available
}

//func (l *Link) getItems() map[string]int {
//	return l.From.Recipe.Product
//}
