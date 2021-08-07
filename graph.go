package imedge

type Graph struct {
	positiveVertices []uint
	negativeVertices []uint
	edges []uint
}

func NewGraph() *Graph {
	return &Graph{
		edges: []uint{0: 0, 1: 0},
	}
}

func (g *Graph) Create() uint {
	id := uint(len(g.positiveVertices))
	g.positiveVertices = append(g.positiveVertices, 0) // TODO: guard with mutex
	g.negativeVertices = append(g.negativeVertices, 0) // TODO: guard with mutex
	return id
}

func (g *Graph) ReadPositive(tail uint) []uint {
	positives := make([]uint, 0)
	edge := g.positiveVertices[tail]
	if edge == 0 {
		return positives
	}

	positives = append(positives, g.edges[edge])

	for next := g.edges[edge + 1]; next != 0; next = g.edges[edge + 1] {
		edge = g.edges[next]
		positives = append(positives, g.edges[edge])
	}

	return positives
}

func (g *Graph) ReadNegative(head uint) []uint {
	positives := make([]uint, 0)
	edge := g.negativeVertices[head]
	if edge == 0 {
		return positives
	}

	positives = append(positives, g.edges[edge])

	for next := g.edges[edge + 1]; next != 0; next = g.edges[edge + 1] {
		edge = g.edges[next]
		positives = append(positives, g.edges[edge])
	}

	return positives
}

func (g *Graph) Connect(tail uint, head uint) {

	positive := uint(len(g.edges))
	g.edges = append(g.edges, head) // TODO: use immutable data structure
	g.edges = append(g.edges, g.positiveVertices[tail]) // TODO: use immutable data structure
	g.positiveVertices[tail] = positive // TODO: guard with mutex

	negative := uint(len(g.edges))
	g.edges = append(g.edges, tail) // TODO: use immutable data structure
	g.edges = append(g.edges, g.negativeVertices[head]) // TODO: use immutable data structure
	g.negativeVertices[head] = negative // TODO: guard with mutex
}
