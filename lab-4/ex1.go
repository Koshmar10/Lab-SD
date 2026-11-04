package main

import "sync"

type Nod struct {
	id     int
	left   chan int
	right  chan int
	status string
}

func (n *Nod) transform(rounds int, wg *sync.WaitGroup) {
	defer wg.Done()
	send_id := n.id
	for r := 0; r < rounds; r++ {
		n.right <- send_id
		recv_id := <-n.left

		if recv_id == n.id {
			n.status = "leader"
			println("Sunt liderul", n.id)
			send_id = -1
		} else if recv_id > n.id {
			send_id = recv_id
		} else {
			send_id = -1
		}
	}
}

func main() {
	var wg sync.WaitGroup
	nodes := []Nod{}

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)
	ch4 := make(chan int, 1)
	ch5 := make(chan int, 1)

	nodes = append(nodes, Nod{1, ch5, ch1, "non-lider"})
	nodes = append(nodes, Nod{2, ch1, ch2, "non-lider"})
	nodes = append(nodes, Nod{3, ch2, ch3, "non-lider"})
	nodes = append(nodes, Nod{4, ch3, ch4, "non-lider"})
	nodes = append(nodes, Nod{5, ch4, ch5, "non-lider"})

	for i := range nodes {
		wg.Add(1)
		go nodes[i].transform(len(nodes), &wg)
	}
	wg.Wait()
}
