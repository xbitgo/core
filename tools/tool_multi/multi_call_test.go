package tool_multi

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestGroupCall(t *testing.T) {
	funcList := make([]func(ctx context.Context) (interface{}, error), 0)
	for i := 0; i < 5; i++ {
		funcList = append(funcList, func(ctx context.Context) (interface{}, error) {
			x := rand.Int31n(9)
			time.Sleep(time.Duration(x) * time.Second)

			return fmt.Sprintf("rs=%d", x), nil
		})
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rs := MultiCall(ctx, 5*time.Second, funcList...)
	for i, r := range rs {
		fmt.Println(r.Val)
		fmt.Println(i, r.Val, r.Err, r.Time.IsZero())
	}
	select {
	case <-time.After(20 * time.Second):
		for i, r := range rs {
			fmt.Println(r.Val)
			fmt.Println(i, r.Val, r.Err, r.Time)
		}
	}
}
