package main

import (
	"container/list"
	"fmt"
)

type Vertex struct {
	data        interface{}
	inputEdges  *list.List
	outputEdges *list.List
}

func (v *Vertex) create(value interface{}) {
	v.data = value
	v.inputEdges = list.New()
	v.outputEdges = list.New()
}

func (v *Vertex) addInputEdge(e *Edge) {
	v.inputEdges.PushBack(e)
}

func (v *Vertex) addOutputEdge(e *Edge) {
	v.outputEdges.PushBack(e)
}

type Edge struct {
	weight float64
	start  *Vertex
	end    *Vertex
}

func (e *Edge) create(w float64, start, end *Vertex) {
	e.weight = w
	e.start = start
	e.end = end
}

type Graph struct {
	vertex *list.List
	edges  *list.List
}

func (g *Graph) create() {
	g.vertex = list.New()
	g.edges = list.New()
}

func (g *Graph) addVertex(data interface{}) {
	var vertex *Vertex = &Vertex{}
	vertex.create(data)

	g.vertex.PushBack(vertex)
}

func (g *Graph) addEdge(weight float64, dataStart, dataEnd interface{}) {
	var start *Vertex = g.getVertex(dataStart)
	var end *Vertex = g.getVertex(dataEnd)

	var edge *Edge = &Edge{}
	edge.create(weight, start, end)

	start.addOutputEdge(edge)
	end.addInputEdge(edge)

	g.edges.PushBack(edge)
}

func (g *Graph) getVertex(data interface{}) *Vertex {
	var vertex *Vertex = nil

	for l := g.vertex.Front(); l != nil; l = l.Next() {
		v := l.Value.(*Vertex)

		if v.data == data {
			vertex = l.Value.(*Vertex)
			break
		}
	}

	return vertex
}

func (g *Graph) BFS() {
	visiteds := list.New()
	queue := list.New()

	element := g.vertex.Front()
	current := element.Value.(*Vertex)

	// current := g.getVertex("Rogerin")

	visiteds.PushBack(current)
	fmt.Println(current.data)

	queue.PushBack(current)

	for queue.Len() > 0 {
		var visited *Vertex = queue.Front().Value.(*Vertex)

		for e := visited.outputEdges.Front(); e != nil; e = e.Next() {
			var next *Vertex = e.Value.(*Edge).end

			if vis := isVisited(visiteds, next); !vis {
				visiteds.PushBack(next)
				fmt.Println(next.data)
				queue.PushBack(next)
			}

		}

		queue.Remove(queue.Front())
	}
}

func isVisited(l *list.List, e *Vertex) bool {
	for i := l.Front(); i != nil; i = i.Next() {
		var vertex *Vertex = i.Value.(*Vertex)
		if vertex == e {
			return true
		}
	}

	return false
}

func main() {
	var graph *Graph = &Graph{}
	graph.create()

	graph.addVertex("Guilherme")
	graph.addVertex("Yasmim")
	graph.addVertex("João")
	graph.addVertex("Clebin")
	graph.addVertex("Rogerin")

	graph.addEdge(2, "Guilherme", "Yasmim")
	graph.addEdge(3, "Yasmim", "Clebin")
	graph.addEdge(1, "Clebin", "João")
	graph.addEdge(1, "Guilherme", "João")
	graph.addEdge(2, "Rogerin", "Yasmim")
	graph.addEdge(3, "Rogerin", "Guilherme")

	graph.BFS()
}
