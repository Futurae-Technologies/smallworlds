package graph

// Graph is an interface that all generated graphs
// (from the subpackages) satisfy.
type Graph interface {
	Edges() []Edge
	Nodes() []Node
}

// Edge is an interface representing a _directed_ graph edge.
// Generated graphs provide edges in both directions
// if they generate undirected graphs.
type Edge interface {
	From() Node
	To() Node
}

// Node is an interface that has a String representation, and
// nodes are equivalent if they have the same String representations.
type Node interface {
	String() string
}

// Contains returns true if the nodes contain node.
func Contains(nodes []Node, node Node) bool {
	for _, i := range nodes {
		if node.String() == i.String() {
			return true
		}
	}
	return false
}

// D3Json function maps a Graph into a JSON object that
// D3 JavaScript library can display.
//
//	Schema ::= {
//		"nodes": [{"id":"ID", "group": 1}],
//		"links":[{"source":"ID", "target":"ID", "value":1}]
//	}
func D3Json(w Graph) map[string]interface{} {
	nodesObj := make([]interface{}, 0, 0)

	for _, node := range w.Nodes() {
		nodeObj := make(map[string]interface{})
		nodeObj["id"] = node.String()
		nodeObj["group"] = 1 // will be cluster id

		nodesObj = append(nodesObj, nodeObj)
	}

	linksObj := make([]interface{}, 0, 0)

	for _, e := range w.Edges() {
		edgeObj := make(map[string]interface{})
		edgeObj["source"] = e.From().String()
		edgeObj["target"] = e.To().String()
		edgeObj["value"] = 1

		linksObj = append(linksObj, edgeObj)
	}

	body := make(map[string]interface{})
	body["nodes"] = nodesObj
	body["links"] = linksObj
	return body
}
