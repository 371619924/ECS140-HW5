package lgraph

type node uint

type edge struct {
	destination node
	label       rune
}

// LGraph is a function representing a directed labeled graph. If the node exists
// in the graph, the function returns true along with the set of outgoing edges
// from that node, otherwise false and nil.
type LGraph func(node) ([]edge, bool)

// FindSequence returns (S, true) if there is a sequence S of length k from node
// s to node t in graph g1 and S is not a sequence from s to t in graph g2; else
// it returns (nil, false).
func FindSequence(g1, g2 LGraph, s, t node, k uint) ([]rune, bool) {
	// TODO: Complete the function.
	sequences := make(chan []rune)
	go func() {
		defer close(sequences)
		findSequences(g1, s, t, k, []rune{}, sequences)

}()


var answer []rune
	found := false
	for sequence := range sequences {
		if !hasSequence(g2, s, t, sequence) && !found {
			answer = sequence
			found = true
		}
	}

	if found {
		return answer, true
	}
	return nil, false
}

func findSequences(g LGraph, current, target node, length uint, path []rune, out chan<- []rune) {
	if length == 0 {
		if current == target {
			sequence := make([]rune, len(path))
			copy(sequence, path)
			out <- sequence
		}
		return
	}

	edges, exists := g(current)
	if !exists {
		return
	}

	for _, next := range edges {
		path = append(path, next.label)
		findSequences(g, next.destination, target, length-1, path, out)
		path = path[:len(path)-1]
	}
}

func hasSequence(g LGraph, current, target node, sequence []rune) bool {
	_, exists := g(current)
	if !exists {
		return false
	}

	if len(sequence) == 0 {
		return current == target
	}

	edges, _ := g(current)
	for _, next := range edges {
		if next.label == sequence[0] && hasSequence(g, next.destination, target, sequence[1:]) {
			return true
		}
	}
	return false
}