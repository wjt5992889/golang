package main

import(
	"time"
	"fmt"
)



func Producer(data chan<-int){
	i :=1
	for{
		select {
		case data <- i:
			fmt.Println("produce")
			i++
			time.Sleep(time.Second * 1)

		}
	}

}

func consumer(data chan int){
	
	for{
		select {
		case <-data:
			fmt.Println("take" )
		}
	}

}

func main() {
	data := make(chan int, 10)
	
	
	go Producer(data)
	time.Sleep(time.Second * 10)
	go consumer(data)
	time.Sleep(time.Second * 11)
}