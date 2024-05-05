package point

type Point struct {
	Name      string
	Latitude  float64
	Longitude float64
}

var Points = []Point{
	{"La Kave Maltée", 45.7654056, 4.8529139},
	{"Béguin, Bar à manger", 45.7686443, 4.8509891},
	{"Le Regain", 45.7671048, 4.8276292},
}
