# Graphs

Um Graph ou Grafo é uma estrutura de dados que representa um par de conjuntos denominados:

- `Vértices` (Vertex)

- `Arestas` (Edges)

Cada `Aresta` está associoada a dois `Vértices` sendo o primeiro o ponto inicial e o segundo o ponto final da aresta.

Se tivermos uma aresta `a` e esta tiver um ponto inicial `v` e um ponto final `w` dizemos que `a` vai de `v` até `w`.

Um exemplo na prática é que podemos representar a estrutura da rede WWW usando grafo onde os `vértices` são as páginas web e as `arestas` são os links que direcionam para as outras páginas. 

Abaixo temos imagens de grafos para um melhor entendimento.

![graph](https://user-images.githubusercontent.com/48635609/102935498-08912080-4485-11eb-9629-8f4bdeb6c93a.png)

![directed](https://user-images.githubusercontent.com/48635609/102935479-fd3df500-4484-11eb-9f17-7cca43e72749.png)

Podemos notar dois grafos com `arestas` diferentes, no primeiro grafo temos uma relação simétrica ou seja se o vértice `v` está associoado a `w` então `w` está associado a `v`, já no segundo temos um grafo direcionado ou também chamado de `digrafo` onde `w` pode está associado a `v` mas `v` não precisa está associado a `w`. 
