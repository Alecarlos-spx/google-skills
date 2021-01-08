# A*

O algoritmo A* é um algoritmo usado para busca de caminho em um [grafo](https://github.com/GuilhermehChaves/google-skills/tree/develop/estrutura_de_dados/graphs). Ele busca o melhor caminho entre um vértice inicial e um vértice final.

Este algoritmo é muito utilizado em jogos, encontrar rotas como em GPS, entre outros...

## Como funciona?

O que o algoritmo faz basicamente é, a partir do vértice atual (Vértice inicial) ele pega todos os seus vizinhos e aplica a função de avaliação `F(n) = G(n) + H(n)` onde: 

- G(n) = Custo de chegada até o vértice
- H(n) = Valor heurístico do vértice - Indica a distância do vértice até o vértice final

O vértice que tiver como retorno o menor valor de F(n) passará a ser o vértice atual, este procedimento se repete até que o vértice atual seja o vértice final.

## Exemplo

![astar](https://user-images.githubusercontent.com/48635609/104031436-7c346c80-51ab-11eb-8c5f-6e432b5b1146.PNG)

Levando em consideração o digrafo acima, Queremos partir do ponto A ao G, onde já temos os valores de heurítica definidos como:

| Vétice | H |
| ------ | --- |
| A | 30 |
| B | 26 |
| C | 21 |
| D | 7 |
| E | 22 |
| F | 36 |
| G | 0 |

Lembrando que a função de heurística pode ser calculada no próprio algoritmo, mas neste exemplo já teremos ela definida.

