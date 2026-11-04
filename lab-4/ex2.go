package main

import "sync"

type Nod struct {
	id     int
	in     []chan int
	out    []chan int
	status string
}

func (n *Nod) transform(rounds int, wg *sync.WaitGroup) {
	defer wg.Done()
	max_id := n.id
	for r := 0; r < rounds; r++ {
		for _, ch := range n.out {
			ch <- max_id
		}
		for _, ch := range n.in {
			recv_id := <-ch
			if recv_id > max_id {
				max_id = recv_id
			}
		}
	}

	if max_id == n.id {
		n.status = "lider"
		println("Sunt liderul", n.id)
	} else {
		println("Nodul", n.id, "stie ca liderul este", max_id)
	}
}

func main() {
	var wg sync.WaitGroup
	nodes := []Nod{}

	nodes = append(nodes, Nod{4, nil, nil, "non-lider"})
	nodes = append(nodes, Nod{9, nil, nil, "non-lider"})
	nodes = append(nodes, Nod{2, nil, nil, "non-lider"})
	nodes = append(nodes, Nod{7, nil, nil, "non-lider"})
	nodes = append(nodes, Nod{5, nil, nil, "non-lider"})

	edges := [][2]int{{0, 1}, {1, 2}, {2, 0}, {2, 3}, {3, 4}, {4, 1}, {0, 3}}
	for _, e := range edges {
		ch := make(chan int, 1)
		nodes[e[0]].out = append(nodes[e[0]].out, ch)
		nodes[e[1]].in = append(nodes[e[1]].in, ch)
	}

	for i := range nodes {
		wg.Add(1)
		go nodes[i].transform(len(nodes), &wg)
	}
	wg.Wait()
}
