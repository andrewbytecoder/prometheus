package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	ctx := context.Background()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			return
		//	因为Tick 每次都创建一个新的计时器，所以 ticker 会被触发很多次
		// 这种case下不能有default 要阻塞等待否则会一直触发default
		case <-time.Tick(1 * time.Second):
			fmt.Println("tick")
		default: // 有没有default 决定是否进行阻塞等待
			fmt.Println("default")
		}
	}

}
