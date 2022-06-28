package tool_multi

import (
	"context"
	"github.com/pkg/errors"
	"sync"
	"time"

	"github.com/xbitgo/core/log"
)

var TimeoutErr = errors.New("timeout error")

type Result struct {
	Val  interface{}
	Err  error
	Time time.Time
}

func MultiCall(ctx context.Context, timeout time.Duration, fun ...func(ctx context.Context) (interface{}, error)) []Result {
	var (
		wg               = sync.WaitGroup{}
		rs               = make([]Result, len(fun))
		done             = make(chan struct{}, 0)
		deadlineExceeded = false
	)
	for i, f := range fun {
		wg.Add(1)
		go func(idx int, fu func(ctx context.Context) (interface{}, error)) {
			rs[idx] = Result{
				Err: TimeoutErr,
			}
			r, err := fu(ctx)
			rs[idx] = Result{
				Val:  r,
				Err:  err,
				Time: time.Now(),
			}
			wg.Done()
		}(i, f)
	}
	go func() {
		wg.Wait()
		if !deadlineExceeded {
			done <- struct{}{}
		}
	}()

	select {
	case <-done:
	case <-ctx.Done():
		log.With().TraceID(ctx).Errorf("MultiCall timeout on ctx cancel")
		deadlineExceeded = true
		close(done)
	case <-time.After(timeout):
		log.With().TraceID(ctx).Errorf("MultiCall timeout[%v]", timeout)
		deadlineExceeded = true
		close(done)
	}
	return rs
}
