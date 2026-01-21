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
	ctx, cancel := context.WithCancel(context.Background())

	correctNumber := 5

	prlWorkers := workers.New().SetWorkerCount(10)

	prlWorkers = prlWorkers.SetTask(func() error {
		for {
			select {
			case <-ctx.Done():
				return nil
			default:
				time.Sleep(1 * time.Second)
				guess := rand.Intn(10)
				fmt.Println("GUESSING:", guess)
				if guess == correctNumber {
					cancel()
				}
			}
		}
	})
	err := prlWorkers.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}

	// text := "These pretzels are making me thirsty."
	//	// Approach 1: Using md5.Sum() for simple, one-off hashing
	//	hasherBytes := md5.Sum([]byte(text))
	//	// Convert the [16]byte array to a hexadecimal string
	//	hashString := hex.EncodeToString(hasherBytes[:])
	//	fmt.Printf("Hash 1: %s\n", hashString)
}
