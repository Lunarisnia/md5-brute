package brute

import (
	"fmt"
	"testing"
)

func Test_Crack(t *testing.T) {
	t.Run("crack successfull all lowercase", func(t *testing.T) {
		bruteForcer := New()

		bruteForcer = bruteForcer.SetGoalTest(func(guess string) bool {
			return guess == "foo"
		}).SetTextLength(3)

		result, err := bruteForcer.Crack()
		if err != nil {
			t.Error(err)
		}
		fmt.Println(result)
	})
}
