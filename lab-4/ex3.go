package main

import (
	"fmt"
	"sync"
)

type Nod struct {
	id    int
	value float64
	in    []chan float64
	out   []chan float64
}

func (n *Nod) transform(rounds int, wg *sync.WaitGroup) {
	defer wg.Done()
	for r := 0; r < rounds; r++ {
		for _, ch := range n.out {
			ch <- n.value
		}
		sum := n.value
		for _, ch := range n.in {
			sum += <-ch
		}
		n.value = sum / float64(len(n.in)+1)
	}

	decision := 0
	if n.value > 0.5 {
		decision = 1
	}
	fmt.Printf("Nodul %d: media %.4f -> decizia %d\n", n.id, n.value, decision)
}

func main() {
	var wg sync.WaitGroup
	nodes := []Nod{}
	rounds := 20

	nodes = append(nodes, Nod{1, 1, nil, nil})
	nodes = append(nodes, Nod{2, 0, nil, nil})
	nodes = append(nodes, Nod{3, 1, nil, nil})
	nodes = append(nodes, Nod{4, 1, nil, nil})
	nodes = append(nodes, Nod{5, 0, nil, nil})

	for _, n := range nodes {
		fmt.Println("Nodul", n.id, "voteaza", n.value)
	}

	edges := [][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 0}}
	for _, e := range edges {
		ch1 := make(chan float64, 1)
		ch2 := make(chan float64, 1)
		nodes[e[0]].out = append(nodes[e[0]].out, ch1)
		nodes[e[1]].in = append(nodes[e[1]].in, ch1)
		nodes[e[1]].out = append(nodes[e[1]].out, ch2)
		nodes[e[0]].in = append(nodes[e[0]].in, ch2)
	}

	for i := range nodes {
		wg.Add(1)
		go nodes[i].transform(rounds, &wg)
	}
	wg.Wait()
}
