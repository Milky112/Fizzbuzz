package usecase

import (
	"context"
	"fizzbuzz/model"
	"strconv"
	"time"
)

func FizzbuzzCalculate(ctx context.Context, req model.FizzBuzzRequest) []string {
	var result []string
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for i := req.From; i <= req.To; i++ {
		if i%15 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return result
}

func StopGraceful() {
	select {
	case <-time.After(10 * time.Second):
		return
	}
}
