package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Punct struct {
	x float64
	y float64
}

type Robot struct {
	id  int32
	poz Punct
}

func (r *Robot) update(new_poz Punct) {
	r.poz.x = new_poz.x
	r.poz.y = new_poz.y
}

func check_conv(roboti []Robot) bool {
	for i := 0; i < len(roboti)-1; i++ {
		if dist(roboti[i].poz, roboti[i+1].poz) > 1e-6 {
			return false
		}
	}
	return true
}

func dist(a Punct, b Punct) float64 {
	f1 := a.x - b.x
	f2 := b.y - a.y
	return math.Sqrt(f1*f1 + f2*f2)
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	nr_roboti, _ := strconv.Atoi(scanner.Text())

	next_id := int32(1)
	roboti := []Robot{}
	for scanner.Scan() {
		coords := strings.Fields(scanner.Text())
		x, _ := strconv.ParseFloat(coords[0], 64)
		y, _ := strconv.ParseFloat(coords[1], 64)
		roboti = append(roboti, Robot{id: next_id, poz: Punct{x, y}})
		next_id += 1
	}

	fmt.Println("Numar roboti:", nr_roboti)
	raza := 16.0
	runda := 0
	for !check_conv(roboti) && runda < 1000 {
		new_pozitii := []Punct{}
		for i := 0; i < len(roboti); i++ {
			r1 := roboti[i]
			vecini := 1
			punct := r1.poz
			for j := 0; j < len(roboti); j++ {
				r2 := roboti[j]
				if dist(r1.poz, r2.poz) < raza && r1.id != r2.id {
					vecini += 1
					punct.x += r2.poz.x
					punct.y += r2.poz.y
				}
			}
			new_poz := Punct{x: punct.x / float64(vecini), y: punct.y / float64(vecini)}
			new_pozitii = append(new_pozitii, new_poz)
		}
		for i := range roboti {
			roboti[i].update(new_pozitii[i])
		}
		runda += 1
	}
	fmt.Println("Runde:", runda)
	for _, r := range roboti {
		fmt.Println("Robot", r.id, ":", r.poz.x, r.poz.y)
	}
}
