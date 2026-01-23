package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Lunarisnia/md5-brute/md5-go/internal/brute"
	"github.com/Lunarisnia/md5-brute/md5-go/internal/hasher"
	"github.com/Lunarisnia/md5-brute/md5-go/internal/workers"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	rawText := "hell"
	hashedText := "4229d691b07b13341da53f17ab9f2416" // hell
	// hashedText := "4124bc0a9335c27f086f24ba207a4912" // aa

	totalWorker := uint(2)
	prlWorkers := workers.New().SetWorkerCount(totalWorker)
	prlWorkers = prlWorkers.SetTask(func(id uint) error {
		// NOTE: 32 - 50 (W 0)
		startRune := rune(32 + (id * (94 / totalWorker)))
		endRune := rune(uint(startRune) + (94 / totalWorker))
		bruteForcer := brute.New().SetTextLength(uint(len(rawText))).SetStartRune(startRune).SetEndRune(endRune)
		bruteForcer = bruteForcer.SetGoalTest(func(guess string) bool {
			return hasher.MD5(guess) == hashedText
		})

		// FIX: I think the context should be passed to crack and it is the one that listen on <-ctx.Done

		result, err := bruteForcer.Crack(ctx)
		if err != nil {
			return nil
		}
		fmt.Println("RESULT:", result)
		cancel()
		return nil
	})
	err := prlWorkers.Run(ctx)
	if err != nil {
		log.Fatalln(err)
	}
}
