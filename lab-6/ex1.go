package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Info struct {
	ID    int
	Value int
}

type Memory map[Info]struct{}

type Node struct {
	ID        int
	Neighbors []int
	M         Memory
	crashed   bool
}

func diametru(adj [][]int) int {
	n := len(adj)
	diam := 0
	for src := 0; src < n; src++ {
		dist := make([]int, n)
		for i := range dist {
			dist[i] = -1
		}
		dist[src] = 0
		queue := []int{src}
		for len(queue) > 0 {
			u := queue[0]
			queue = queue[1:]
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					queue = append(queue, v)
				}
			}
		}
		for _, d := range dist {
			if d == -1 {
				return -1
			}
			if d > diam {
				diam = d
			}
		}
	}
	return diam
}

func copie(m Memory) Memory {
	rez := Memory{}
	for info := range m {
		rez[info] = struct{}{}
	}
	return rez
}

func floodSet(noduri []Node, s int, diam int) int {
	nr_crash := 0
	t := 0
	for {
		crash_acum := -1
		if nr_crash < s && rand.IntN(2) == 0 {
			corecte := []int{}
			for _, nod := range noduri {
				if !nod.crashed {
					corecte = append(corecte, nod.ID)
				}
			}
			if len(corecte) > 1 {
				crash_acum = corecte[rand.IntN(len(corecte))]
				nr_crash += 1
			}
		}

		mesaje := make([][]Memory, len(noduri))
		for i := range noduri {
			if noduri[i].crashed {
				continue
			}
			trimis := copie(noduri[i].M)
			vecini := noduri[i].Neighbors
			if i == crash_acum {
				vecini = vecini[:rand.IntN(len(vecini)+1)]
			}
			for _, j := range vecini {
				mesaje[j] = append(mesaje[j], trimis)
			}
		}
		if crash_acum != -1 {
			noduri[crash_acum].crashed = true
		}

		for i := range noduri {
			if noduri[i].crashed {
				continue
			}
			for _, u := range mesaje[i] {
				for info := range u {
					noduri[i].M[info] = struct{}{}
				}
			}
		}

		if t > (s+1)*diam {
			return t
		}
		t += 1
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	prima := strings.Fields(scanner.Text())
	n, _ := strconv.Atoi(prima[0])
	s, _ := strconv.Atoi(prima[1])

	scanner.Scan()
	tokenuri := strings.Fields(scanner.Text())

	noduri := []Node{}
	adj := [][]int{}
	for i := 0; i < n; i++ {
		scanner.Scan()
		vecini := []int{}
		for _, v := range strings.Fields(scanner.Text()) {
			vecin, _ := strconv.Atoi(v)
			vecini = append(vecini, vecin)
		}
		adj = append(adj, vecini)
		valoare, _ := strconv.Atoi(tokenuri[i])
		noduri = append(noduri, Node{
			ID:        i,
			Neighbors: vecini,
			M:         Memory{Info{ID: i, Value: valoare}: struct{}{}},
		})
	}

	diam := diametru(adj)
	if diam == -1 {
		fmt.Println("graful nu este tare conex")
		return
	}
	t_final := floodSet(noduri, s, diam)

	for _, nod := range noduri {
		if nod.crashed {
			continue
		}
		memorie := []Info{}
		for info := range nod.M {
			memorie = append(memorie, info)
		}
		slices.SortFunc(memorie, func(a, b Info) int { return a.ID - b.ID })
		perechi := []string{}
		for _, info := range memorie {
			perechi = append(perechi, fmt.Sprintf("(%d, %d)", info.ID, info.Value))
		}
		fmt.Printf("diametru: %d\niteratia de oprire: %d\n%s\n", diam, t_final, strings.Join(perechi, " "))
		return
	}
}
