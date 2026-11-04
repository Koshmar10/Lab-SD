package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
)

type LamportClock struct{ counter int }

func (l *LamportClock) Tick() { l.counter++ }
func (l *LamportClock) Send() { l.counter++ }
func (l *LamportClock) Receive(ts int) {
	l.counter = max(l.counter, ts) + 1
}

type Message struct{ sender, ts int }
type Event struct {
	process          int
	internal_counter int
	kind             string
}

func less(a, b Event) bool {
	if a.internal_counter != b.internal_counter {
		return a.internal_counter < b.internal_counter
	}
	return a.process < b.process
}

func process(id int, rounds int, next chan<- Message, prev <-chan Message, events chan<- Event, wg *sync.WaitGroup) {
	defer wg.Done()
	clock := LamportClock{0}

	for r := 0; r < rounds; r++ {
		internal := rand.Intn(3)
		for i := 0; i < internal; i++ {
			clock.Tick()
			events <- Event{id, clock.counter, "internal"}
		}

		clock.Send()
		next <- Message{id, clock.counter}
		events <- Event{id, clock.counter, "send"}

		msg := <-prev
		clock.Receive(msg.ts)
		events <- Event{id, clock.counter, "recv"}
	}
}

func main() {
	var wg sync.WaitGroup
	events := make(chan Event)
	all_events := []Event{}
	rounds := 3

	ch1 := make(chan Message, 1)
	ch2 := make(chan Message, 1)
	ch3 := make(chan Message, 1)

	wg.Add(3)
	go process(1, rounds, ch1, ch3, events, &wg)
	go process(2, rounds, ch2, ch1, events, &wg)
	go process(3, rounds, ch3, ch2, events, &wg)

	go func() {
		wg.Wait()
		close(events)
	}()

	for e := range events {
		all_events = append(all_events, e)
	}

	sort.Slice(all_events, func(i, j int) bool {
		return less(all_events[i], all_events[j])
	})

	fmt.Println("Evenimente ordonate crescator:")
	for _, e := range all_events {
		fmt.Printf("%3d  P%d  %s\n", e.internal_counter, e.process, e.kind)
	}
}
