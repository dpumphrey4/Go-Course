package gopherstest

import (
	"fmt"
	"math/rand"
	"sync"
)

type gopher struct {
	Name   string
	Age    uint16
	Number uint16
}

// Gopher names
var names = []string{"Aria", "Nolan", "Sienna", "Jasper", "Kiara", "Theo"}

// This is running on the main goroutine
func MainGopher() {
	fmt.Println("Basic Gophers============================================")

	// Making a waitgroup
	var wg sync.WaitGroup

	// Making the channel
	gophers := make(chan gopher, 20)

	// Starting the goroutines
	wg.Go(func() { gopherSleep(gophers) })
	wg.Go(func() { gopherKill(gophers) })
	wg.Go(func() { gopherRun(gophers) })

	// Filling the channel with 100 gophers
	for i := 0; i < 100; i++ {
		newGopher := gopher{
			Name:   names[rand.Intn(6)],
			Age:    uint16(rand.Intn(100)),
			Number: uint16(i),
		}

		gophers <- newGopher
	}

	// Closing the channel
	close(gophers)

	// Waiting until completion
	wg.Wait()
}

func gopherSleep(gophers chan gopher) {
	for {
		localGopher, ok := <-gophers
		if ok == false {
			break
		}
		fmt.Printf("Number: %d, Gopher: %s, age: %d, is sleeping.\n", localGopher.Number, localGopher.Name, localGopher.Age)
	}

}

func gopherKill(gophers chan gopher) {
	for {
		localGopher, ok := <-gophers
		if ok == false {
			break
		}
		fmt.Printf("Number: %d, Gopher: %s, age: %d, is dead.\n", localGopher.Number, localGopher.Name, localGopher.Age)
	}
}

func gopherRun(gophers chan gopher) {
	for {
		localGopher, ok := <-gophers
		if ok == false {
			break
		}
		fmt.Printf("Number: %d, Gopher: %s, age: %d, is dead.\n", localGopher.Number, localGopher.Name, localGopher.Age)
	}
}
