package main

func sender(n int, c chan int) {
	c <- 10
}

func main() {
	c := make(chan int)
	go sender(10, c)
	<-c
}
