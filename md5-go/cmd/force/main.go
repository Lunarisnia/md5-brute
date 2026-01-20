package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/Lunarisnia/md5-brute/md5-go/internal/workers"
)

func main() {
	ctx := context.Background()

	correctNumber := 37

	prlWorkers := workers.New()

	counter := make(chan int)
	found := make(chan bool)
	prlWorkers = prlWorkers.SetTask(func() error {
		for {
			time.Sleep(1 * time.Second)
			counter <- rand.Intn(100)
			fmt.Println("Guessing")
			if <-found {
				break
			}
		}

		return nil
	})
	prlWorkers = prlWorkers.SetWorkerCount(3)
	go err := prlWorkers.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// wCounter := 0
	// for {
	// 	<-counter
	// 	wCounter++
	// 	fmt.Println("Task Finished!")
	// 	if wCounter == 3 {
	// 		os.Exit(0)
	// 	}
	// }

	// text := "These pretzels are making me thirsty."
	//
	//	// Approach 1: Using md5.Sum() for simple, one-off hashing
	//	hasherBytes := md5.Sum([]byte(text))
	//	// Convert the [16]byte array to a hexadecimal string
	//	hashString := hex.EncodeToString(hasherBytes[:])
	//	fmt.Printf("Hash 1: %s\n", hashString)
}
