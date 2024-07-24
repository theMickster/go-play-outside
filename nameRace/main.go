package main

import (
	"fmt"
	"sync"
)

func main() {
	people := []string{"Olivia", "Noah", "Amelia", "Liam", "Charlotte", "Michael", "Isabella", "Muhammad", "Maverick", "Sebastian", "Scarlet", "Penelope", "Delilah"}
	peopleMap := make(map[string]int)

	var myWaitGroup sync.WaitGroup
	myWaitGroup.Add(len(people))

	var myMutex sync.Mutex

	for _, name := range people {
		name := name
		go func(babyName string) {
			defer myWaitGroup.Done()
			myMutex.Lock()
			defer myMutex.Unlock()
			peopleMap[name] = len(babyName)
		}(name)
	}

	myWaitGroup.Wait()

	fmt.Println(peopleMap)
}
