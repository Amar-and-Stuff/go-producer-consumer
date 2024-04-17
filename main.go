package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"sync"
	"time"
)

func main() {
	mu := sync.Mutex{}
	pcon := sync.NewCond(&mu)
	ccon := sync.NewCond(&mu)
	noOfProducers := 3000
	noOfConsumers := 3000
	market := 0
	capacity := 30

	for i := 0; i < noOfConsumers; i++ {
		go func(timeToSleep time.Duration) {
			for true {
				ccon.L.Lock()
				if market > 0 {
					market--
					pcon.Broadcast()
				} else {
					ccon.Wait()
				}
				ccon.L.Unlock()
			}
		}(time.Millisecond * time.Duration(rand.Intn(5000)+1000))
	}

	for i := 0; i < noOfProducers; i++ {
		go func(timeToSleep time.Duration) {
			for true {
				pcon.L.Lock()
				if market < capacity {
					market++
					ccon.Broadcast()
				} else {
					pcon.Wait()
				}
				pcon.L.Unlock()
			}
		}(time.Millisecond * time.Duration(rand.Intn(5000)+1000))
	}

	itr := 0
	for true {
		fmt.Printf("%d: ", itr)
		itr++
		for i := 0; i < market; i++ {
			fmt.Print("1 ")
		}
		time.Sleep(time.Millisecond * 50)
		fmt.Println()
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
