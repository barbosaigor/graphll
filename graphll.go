package graphll

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// ErrElementNotFound indicates that no element was found
	ErrElementNotFound = errors.New("Element not found")
)

// Node stores a graph node data
type Node struct {
	Weight uint32
	Deps   set
}

// GraphLL implements a graph data structure
type GraphLL map[string]Node

// New creates a graph
func New() GraphLL {
	return make(GraphLL)
}

// Add attaches a node to the graph
func (g GraphLL) Add(name string, weight uint32, deps []string) {
	if deps == nil {
		deps = []string{}
	}
	if _, ok := g[name]; !ok {
		g[name] = Node{weight, newSet(deps)}
	} else {
		g[name].Deps.union(newSet(deps))
	}
}

// Deps gets the dependencies of a node
func (g GraphLL) Deps(name string) ([]string, error) {
	if _, ok := g[name]; !ok {
		return nil, ErrElementNotFound
	}
	return g[name].Deps.toSlice(), nil
}

// Weight gets the weight of a node
func (g GraphLL) Weight(name string) (uint32, error) {
	if _, ok := g[name]; !ok {
		return 0, ErrElementNotFound
	}
	return g[name].Weight, nil
}

// String creates a string containing informations about the graph
func (g GraphLL) String() string {
	if len(g) == 0 {
		return ""
	}
	var lstr []string
	for name, node := range g {
		lstr = append(lstr, fmt.Sprintf("[%v|%v]: %v", name, node.Weight, node.Deps))
	}
	return strings.Join(lstr, "\n")
}
