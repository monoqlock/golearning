package main

import (
	"fmt"
	"time"
)

/**
時間のかかる同期呼び出しを放棄するには、select文をtime.Afterで使用します。
 */

func main()  {
	c := make(chan error, 1)
	go func() {c <- fmt.Errorf("call error ")}()
	select {
	case err := <- c:
		fmt.Println("use error and reply. err=", err)
	case <-time.After(time.Second * 10):
		fmt.Println("call time out")
	}
}
