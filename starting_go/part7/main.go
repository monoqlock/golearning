package main


import (
	"sync"
	"fmt"
	"time"
)

var st struct{ a, b, c int}


var mutex *sync.Mutex

func UpdateAndPrint(o, i int) {

	mutex.Lock()

	st.a = i
	time.Sleep(time.Millisecond)

	st.b = i
	time.Sleep(time.Millisecond)

	st.c = i
	time.Sleep(time.Millisecond)

	fmt.Println(o, st.a, st.b, st.c)

	mutex.Unlock()
}

func main()  {

	fmt.Println("---start")

	mutex = new(sync.Mutex)

	for i := 0; i < 5; i++ {
		fmt.Println("OUTER: ", i)

		go func() {
			for j:=0; j < 10; j++ {
				UpdateAndPrint(i, j)
			}

		}()
	}

	time.Sleep(5 * time.Second)

	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		wg.Done()
	}()


	wg.Wait()

}