package brute

import (
	"context"
	"fmt"
	"testing"

	"github.com/Lunarisnia/md5-brute/md5-go/internal/hasher"
)

func Test_Crack(t *testing.T) {
	t.Run("crack successfull all lowercase", func(t *testing.T) {
		bruteForcer := New()

		text := "bar"
		hashedText := hasher.MD5(text)
		fmt.Println(hasher.MD5("aa"))
		bruteForcer = bruteForcer.SetGoalTest(func(guess string) bool {
			if hasher.MD5(guess) == hashedText {
				fmt.Println("Original Hash: ", hashedText)
				fmt.Println("Guessed Hash: ", hasher.MD5(guess))
			}
			return hasher.MD5(guess) == hashedText
		}).SetTextLength(uint(len(text)))

		result, err := bruteForcer.Crack(context.Background())
		if err != nil {
			t.Error(err)
		}
		fmt.Println(result)
	})
}
