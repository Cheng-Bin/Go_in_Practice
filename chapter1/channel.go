package chapter1

import "fmt"
import "time"

func printCount(c chan int) {
	num := 0
	for num >= 0 {
		num = <-c
		fmt.Println(num, " ")
	}
}

func printChann() {
	c := make(chan int)
	a := []int{8, 6, 7, 5, 3, 0, -1}

	go printCount(c)

	for _, v := range a {
		c <- v
	}

	time.Sleep(time.Millisecond * 1)
	fmt.Println("End of main")
}
