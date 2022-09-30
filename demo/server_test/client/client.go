package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

func main() {
	ctx, cancelFun := context.WithTimeout(context.Background(), time.Second*8)
	defer cancelFun()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:8887/sleep/1", nil)
	if err != nil {
		panic(err)
	}
	var client = &http.Client{Timeout: 5 * time.Second}
	do, err := client.Do(req)
	if err != nil {
		if netErr, ok := err.(net.Error); ok {
			if netErr.Timeout() {
				fmt.Printf("%v, %v", do, err)
			}
		}
	}

	//spew.Dump(do, err)
}
