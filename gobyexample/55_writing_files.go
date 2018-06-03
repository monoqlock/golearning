package main

import (
	"io/ioutil"
	"os"
	"fmt"
	"bufio"
)
func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngon")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
	check2(err)

	f, err := os.Create("/tmp/dat2")
	check2(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check2(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check2(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()
}

