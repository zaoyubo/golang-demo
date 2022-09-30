package internal

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"golang.org/x/sync/singleflight"
)

type Result string

func find(ctx context.Context, query string) (Result, error) {
	return Result(fmt.Sprintf("result for %q", query)), nil
}

func TestSingleFlight(t *testing.T) {
	var g singleflight.Group
	const n = 5
	waited := int32(n)
	done := make(chan struct{})
	key := "https://weibo.com/1227368500/H3GIgngon"
	for i := 0; i < n; i++ {
		go func(j int) {
			v, _, shared := g.Do(key, func() (interface{}, error) {
				ret, err := find(context.Background(), fmt.Sprint(j))
				return ret, err
			})
			if atomic.AddInt32(&waited, -1) == 0 {
				close(done)
			}
			fmt.Printf("index: %d, val: %v, shared: %v\n", j, v, shared)
		}(i)
	}
	time.Sleep(time.Second * 2)
	<-done

	/*
		一种可能的结果：
		index: 4, val: result for "4", shared: true
		index: 2, val: result for "2", shared: true
		index: 3, val: result for "2", shared: true
		index: 1, val: result for "4", shared: true
		index: 0, val: result for "0", shared: false
	*/

	/*
		一种可能的结果：
		index: 2, val: result for "0", shared: true
		index: 0, val: result for "0", shared: true
		index: 3, val: result for "0", shared: true
		index: 4, val: result for "0", shared: true
		index: 1, val: result for "0", shared: true
	*/
}
