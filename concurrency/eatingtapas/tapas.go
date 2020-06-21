package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
Four friends are having dinner together.
They ordered five dishes to share, each of which consists of between 5 and 10 morsels.

They eat leisurely, spending between 10 seconds and 30 seconds eating each morsel.

Write a program to simulate the meal.
*/
type person string
type dish string

const (
	minMorsel = 5
	maxMorsel = 10
	minEating = 10
	maxEating = 30
)

func (p person) start(food <-chan dish, wg *sync.WaitGroup) {
	fmt.Printf("%s has started eating...\n", p)
	defer wg.Done()
	for f := range food {
		d := getEatingDuration()
		fmt.Printf("%s is enjoying some %s.\n", p, f)
		time.Sleep(d)
		fmt.Printf("%s finished enjoying some %s for %ds.\n", p, f, int(d.Seconds()))
	}
}

func (d dish) start(food chan<- dish, wg *sync.WaitGroup) {
	fmt.Printf("%s has been served...\n", d)
	defer wg.Done()
	for i := 0; i < getMorselsCount(); i++ {
		food <- d
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	persons := []person{person("alice"), person("bob"), person("charlie"), person("david"), person("elise")}
	dishes := []dish{dish("chorizo"), dish("patatas bravas"), dish("prawns"),
		dish("pimientos de padron"), dish("tortilla")}
	// buffer food for amount of morsels equal to persons amount
	food := make(chan dish, len(persons))

	// dishwg to let us know when all morsels are finished being output
	var dishwg sync.WaitGroup
	// personswg to let us know when all persons have finished
	var personwg sync.WaitGroup

	// start the food routines
	dishwg.Add(len(dishes))
	for _, d := range dishes {
		go d.start(food, &dishwg)
	}
	//start the person routines
	personwg.Add(len(persons))
	for _, p := range persons {
		go p.start(food, &personwg)
	}

	// dishes are done
	dishwg.Wait()
	fmt.Println("All the food has finished serving")
	close(food)
	// signal to everyone that the dinner is finished
	// wait for all persons to finish as some persons might be sleeping/eating
	personwg.Wait()
	// finish everything
	fmt.Println("Everyone is finished. Goodbye!")
}

func getMorselsCount() int {
	return rand.Intn(maxMorsel-minMorsel) + minMorsel
}
func getEatingDuration() time.Duration {
	return time.Duration(rand.Intn(maxEating-minEating)+minEating) * time.Second
}
