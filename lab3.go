package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	mess = make(chan int)
)

func Producer(id int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	for {
		a := r1.Intn(20) + 1
		sp := r1.Intn(4) + 1
		fmt.Printf("Producer %d produce %d sleep %d\n", id, a, sp)
		mess <- a
		time.Sleep(time.Duration(sp) * time.Second)
	}
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

func Cusumer(id int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	sp := r1.Intn(4) + 1
	for {
		data, ok := <-mess
		//fmt.Println(ok)
		if ok {
			f := fib(data)
			fmt.Printf("Cusumer %d Fib %d = %d sleep %d\n", id, data, f, sp)
			time.Sleep(time.Duration(sp) * time.Second)
		} else {
			fmt.Printf("Cusumer %d no data\n", id)
		}

	}
}

func main() {
	go Producer(1)
	go Producer(2)
	go Cusumer(1)
	go Cusumer(2)
	go Cusumer(3)
	select {}
}
