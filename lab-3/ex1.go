package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Clock struct{ offset time.Duration }

func (c *Clock) Now() time.Time { return time.Now().Add(c.offset) }

func estimate(t1, t2, t3, t4 time.Time) (theta, delta time.Duration) {
	theta = (t2.Sub(t1) + t3.Sub(t4)) / 2
	delta = t4.Sub(t1) - t3.Sub(t2)
	return theta, delta
}

func measure(c_clock, s_clock *Clock) (theta, delta time.Duration) {

	t1 := c_clock.Now()
	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)

	t2 := s_clock.Now()
	time.Sleep(5*time.Millisecond + time.Duration(rand.Intn(20))*time.Millisecond)

	t3 := s_clock.Now()

	time.Sleep(3*time.Millisecond + time.Duration(rand.Intn(10))*time.Millisecond)
	t4 := c_clock.Now()

	return estimate(t1, t2, t3, t4)
}

func main() {
	initial_offset := 2 * time.Second
	s_clock := &Clock{0}

	// a)
	c_clock := &Clock{initial_offset}
	theta_a, delta_a := measure(c_clock, s_clock)
	c_clock.offset += theta_a
	error_a := c_clock.offset.Abs()

	fmt.Println("theta =", theta_a, "delta =", delta_a)
	fmt.Println("Offset initial:", initial_offset, "-> offset dupa corectie:", c_clock.offset)
	if error_a < initial_offset.Abs() {
		fmt.Println("Offset-ul a scazut")
	} else {
		fmt.Println("Offset-ul NU a scazut")
	}

	// b)
	fmt.Println("\nb) 10 masuratori")
	c_clock = &Clock{initial_offset}
	min_theta, min_delta := measure(c_clock, s_clock)
	fmt.Println("1: theta =", min_theta, "delta =", min_delta)
	for i := 2; i <= 10; i++ {
		theta, delta := measure(c_clock, s_clock)
		fmt.Printf("%d: theta = %v delta = %v\n", i, theta, delta)
		if delta < min_delta {
			min_theta = theta
			min_delta = delta
		}
	}
	c_clock.offset += min_theta
	error_b := c_clock.offset.Abs()

	fmt.Println("theta minim =", min_theta, "delta minim =", min_delta)
	fmt.Println("Offset initial:", initial_offset, "-> offset dupa corectie:", c_clock.offset)

	//a) vs b)
	fmt.Println("\nComparatie")
	fmt.Println("a) delta =", delta_a, "eroare ramasa =", error_a)
	fmt.Println("b) delta =", min_delta, "eroare ramasa =", error_b)
	if error_b < error_a {
		fmt.Println("Varianta b) este mai precisa")
	} else {
		fmt.Println("Varianta a) a fost mai precisa de data asta")
	}
}
