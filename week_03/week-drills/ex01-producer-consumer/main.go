package main


import (
	"fmt"
)


func produce(ch chan int, values []int){

	for _, val := range values {
		ch <- val
	}
	close(ch)
}



func consume(ch chan int){
	
	for val := range ch {
		fmt.Println(val)
	}
}


func main(){

	ch := make(chan int, 5)
	values := []int{1,5,12,27,42}
	go produce(ch, values)
	consume(ch)
}