package main

import (
	"fmt"
	"math/rand/v2"
	"slices"
)

type Message struct {
	SenderID int
	Round    int
	Value    int
}

type Node struct {
	id          int
	x           int
	s           int
	n           int
	isByzantine bool
	inbox       chan Message
}

func (nod *Node) Bcast(sisteme []Node, runda int, valoare int) {
	for j := range sisteme {
		sisteme[j].inbox <- Message{SenderID: nod.id, Round: runda, Value: valoare}
	}
}

func (nod *Node) MaliciousBcastState(sisteme []Node) {
	for j := range sisteme {
		valoare := rand.IntN(2)
		if rand.IntN(2) == 0 {
			valoare = 1 - nod.x
		}
		sisteme[j].inbox <- Message{SenderID: nod.id, Round: 1, Value: valoare}
	}
}

func (nod *Node) MaliciousBcastTie(sisteme []Node) {
	for j := range sisteme {
		sisteme[j].inbox <- Message{SenderID: nod.id, Round: 2, Value: rand.IntN(2)}
	}
}

func (nod *Node) majoritate() (int, int) {
	nr_unu := 0
	for k := 0; k < nod.n; k++ {
		msg := <-nod.inbox
		nr_unu += msg.Value
	}
	if nr_unu >= nod.n/2+1 {
		return 1, nr_unu
	}
	return 0, nod.n - nr_unu
}

func byzFloodSim(n int, s int, valori []int) int {
	//aleg cate noduri malitiase sunt
	indexi_bizantini := []int{}
	for len(indexi_bizantini) < s {
		index := rand.IntN(n)
		if !slices.Contains(indexi_bizantini, index) {
			indexi_bizantini = append(indexi_bizantini, index)
		}
	}
	fmt.Println("noduri care pot fi bizantine:", indexi_bizantini)
	//fac nodurile
	sisteme := []Node{}
	for i := 0; i < n; i++ {
		sisteme = append(sisteme, Node{
			id:          i,
			x:           valori[i],
			s:           s,
			n:           n,
			isByzantine: false,
			inbox:       make(chan Message, n),
		})
	}
	for t := 0; t < s+2; t++ {
		lider := t
		f := rand.IntN(s + 1)
		activi := slices.Clone(indexi_bizantini)
		rand.Shuffle(len(activi), func(a, b int) { activi[a], activi[b] = activi[b], activi[a] })
		activi = activi[:f]
		for i := range sisteme {
			sisteme[i].isByzantine = slices.Contains(activi, i)
		}
		//runda 1
		for i := range sisteme {
			if sisteme[i].isByzantine {
				sisteme[i].MaliciousBcastState(sisteme)
			} else {
				sisteme[i].Bcast(sisteme, 1, sisteme[i].x)
			}
		}
		majoritati := make([]int, n)
		multiplicitati := make([]int, n)
		for i := range sisteme {
			majoritati[i], multiplicitati[i] = sisteme[i].majoritate()
		}
		//runda 2
		if sisteme[lider].isByzantine {
			sisteme[lider].MaliciousBcastTie(sisteme)
		} else {
			sisteme[lider].Bcast(sisteme, 2, majoritati[lider])
		}
		for i := range sisteme {
			tie := (<-sisteme[i].inbox).Value
			if 2*multiplicitati[i] > n+2*s {
				sisteme[i].x = majoritati[i]
			} else {
				sisteme[i].x = tie
			}
		}
		stari := []int{}
		for _, nod := range sisteme {
			stari = append(stari, nod.x)
		}
		fmt.Println("faza", t, "lider", lider, "bizantini activi", activi, "stari", stari)
	}
	//returnez starea primului nod corect
	for _, nod := range sisteme {
		if !slices.Contains(indexi_bizantini, nod.id) {
			return nod.x
		}
	}
	return sisteme[0].x
}

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	valori := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&valori[i])
	}
	start_finala := byzFloodSim(n, s, valori)

	fmt.Println("stare finala = ", start_finala)
}
