package main

import (
	"context"
	"fmt"
	"time"
)

const shortDuration = 500 * time.Microsecond

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancellation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	n := 1
	for {
		select {
		case <-time.After(100 * time.Microsecond):
			fmt.Println("overslept")
			n++
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
