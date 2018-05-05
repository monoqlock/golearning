package funcs

import (
	"fmt"
	"time"
	)

func Channels() {

	var (
		ch0 chan int
		ch1 <-chan int
		ch2 chan<- int
	)

	ch1 = ch0
	ch2 = ch0
	fmt.Printf("ch0=%v, ch1=%v, ch2=%v\n", ch0, ch1, ch2)

	ch3 := make(chan int, 10)
	ch3 <- 5
	i := <-ch3
	fmt.Println(i)

	ch4 := make(chan int, 20 )
	go receiver("1st goroutine", ch4)
	go receiver("2nd goroutine", ch4)
	go receiver("3rd goroutine", ch4)

	i2 := 0
	for i2 < 20 {
		fmt.Printf("Send %d\n",i2)
		ch4 <- i2
		i2++
	}
	close(ch4)

	time.Sleep(3 * time.Second)

}

func receiver(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			break
		}
		fmt.Printf("%s: Receive %d\n", name, i)
	}
	fmt.Println("[" + name + " is done.]")
}

func Select() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	select {
	case <- ch1:
		fmt.Println("receive from ch1")
	case <- ch2:
		fmt.Println("receive from ch2")
	case ch3 <- 3:
		fmt.Println("ch3 received")
	default:
		fmt.Println("default")

	}
}

func SelectMix()  {

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			i := <- ch1
			ch2 <- i * 2
		}
	}()

	go func() {
		for {
			i := <- ch2
			ch3 <- i - 1
		}
	}()

	n := 1
	LOOP:
		for {
			select {
			/* 整数を増分させつつch1に送信 */
			case ch1 <- n:
				n++
			case i := <- ch3:
				fmt.Println("Received ", i)
			default:
				if n > 20 {
					break LOOP
				}
			}
		}
}
