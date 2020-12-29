package main

import (
	"container/list"
	"fmt"
)

type Vertex struct { // struct representando um vértice
	data        interface{} // dado que iremos adidionar no vértice
	inputEdges  *list.List  // arestas de entrada
	outputEdges *list.List  // arestas de saída
	heuristic   float64
	previousG   float64
}

func (v *Vertex) create(value interface{}, h float64) { // função que cria um novo vértice
	v.data = value             // setando o dado do vértice
	v.inputEdges = list.New()  // inicia como uma lista vazia
	v.outputEdges = list.New() // inicia como uma lista vazia
	v.heuristic = h
}

func (v *Vertex) addInputEdge(e *Edge) { // função que adidiona uma aresta de entrada ao grafo
	v.inputEdges.PushBack(e)
}

func (v *Vertex) addOutputEdge(e *Edge) { // função que adiciona uma aresta de saída ao grafo
	v.outputEdges.PushBack(e)
}

type Edge struct { // struct que representa uma aresta
	weight float64 // peso da aresta
	start  *Vertex // vétice de inicio
	end    *Vertex // vértice de final
}

func (e *Edge) create(w float64, start, end *Vertex) { // função que cria um novo vértice
	e.weight = w
	e.start = start
	e.end = end
}

type Graph struct { // struct que representa um grafo
	vertex *list.List // vértices do grafo
	edges  *list.List // arestas do grafo
}

func NewGraph() *Graph {
	return &Graph{
		vertex: list.New(),
		edges:  list.New(),
	}
}

func (g *Graph) addVertex(data interface{}, h float64) *Vertex { // função que adiciona um vértice ao grafo
	var vertex *Vertex = &Vertex{} // define um grafo
	vertex.create(data, h)         // cria um grafo passando o dado
	g.vertex.PushBack(vertex)      // adiciona o vértice na lista de vértices do grafo

	return vertex
}

func (g *Graph) addEdge(weight float64, start, end *Vertex) { // função que adiciona uma aresta entre os grafos
	// var start *Vertex = g.getVertex(dataStart) // Pesquisa o vértice através do dado - inicio
	// var end *Vertex = g.getVertex(dataEnd)     // Pesquisa o vértice através do dado - final

	var edge *Edge = &Edge{}        // define uma nova aresta
	edge.create(weight, start, end) // cria uma nova aresta, passando os vértices de início e final

	start.addOutputEdge(edge) // adiciona a aresta de saida no vértice de ínicio
	end.addInputEdge(edge)    // adiciona a aresta de entrada no vértice de saída
	g.edges.PushBack(edge)    // adiciona a aresta na lista de arestas do grafo
}

func (g *Graph) getVertex(data interface{}) *Vertex { // função que pesquisa um vértice
	var vertex *Vertex = nil // vértice

	for l := g.vertex.Front(); l != nil; l = l.Next() { // loop que percorre cada vértice do grafo
		v := l.Value.(*Vertex) // passando do tipo interface{} para *Vertex

		if v.data == data { // verifica se o dado do vértice atual é igual ao procurado
			vertex = v // vertex é igual ao v
			break      // para o loop
		}
	}

	return vertex // retorna o vértice caso seja encontrado ou nil
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

// --------------------------------- algoritmo A* ----------------------------------------------

type F struct {
	value     float64
	vertex    *Vertex
	previousG float64
}

func (g *Graph) aStar(initialVertex, finalVertex *Vertex) {
	closeds := list.New()
	openeds := list.New()

	closeds.PushBack(initialVertex)

	var father *Vertex = initialVertex
	fmt.Print(father.data.(string))

	for father != finalVertex {
		for i := father.outputEdges.Front(); i != nil; i = i.Next() {
			var edge *Edge = i.Value.(*Edge)
			var children *Vertex = edge.end
			children.previousG = father.previousG + edge.weight

			openeds.PushBack(&F{
				value:     children.previousG + children.heuristic,
				vertex:    children,
				previousG: children.previousG,
			})
		}

		element := getMinF(father, openeds)
		father = element.Value.(*F).vertex
		father.previousG = element.Value.(*F).previousG

		fmt.Print(" -> " + father.data.(string))

		closeds.PushBack(father)
		openeds.Remove(element)
	}
}

func getMinF(old *Vertex, l *list.List) *list.Element {
	var f *F = nil
	var min float64 = 999
	var minF *list.Element = nil

	for i := l.Front(); i != nil; i = i.Next() {
		f = i.Value.(*F)

		if f.value <= min {
			min = f.value
			minF = i
		}
	}

	if minF.Value.(*F).vertex != old {
		return minF
	}

	l.Remove(minF)

	return getMinF(minF.Value.(*F).vertex, l)
}

func main() {
	var graph *Graph = NewGraph()

	a := graph.addVertex("A", 30)
	b := graph.addVertex("B", 26)
	c := graph.addVertex("C", 21)
	d := graph.addVertex("D", 7)
	e := graph.addVertex("E", 22)
	f := graph.addVertex("F", 36)
	g := graph.addVertex("G", 0)

	graph.addEdge(12, a, b)
	graph.addEdge(14, a, c)
	graph.addEdge(9, b, c)
	graph.addEdge(38, b, d)
	graph.addEdge(24, c, d)
	graph.addEdge(7, c, e)
	graph.addEdge(9, d, g)
	graph.addEdge(13, e, d)
	graph.addEdge(29, e, g)
	graph.addEdge(9, e, f)

	graph.aStar(a, g)

}
