package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getPageByte(url string, ch chan []byte) {
	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	ch <- b
}

func main() {
	ch := make(chan []byte)
	go getPageByte("http://www.google.com", ch)
	go getPageByte("http://www.facebook.com", ch)
	go getPageByte("http://www.twitter.com", ch)

	/*
		pageByte := <-ch
		pageByte2 := <-ch
		pageByte3 := <-ch

		fmt.Printf("%s\n", pageByte)
		fmt.Printf("%s\n", pageByte2)
		fmt.Printf("%s\n", pageByte3)
	*/

	//ได้เหมือนกัน
	fmt.Printf("%s\n", <-ch)
	fmt.Printf("%s\n", <-ch)
	fmt.Printf("%s\n", <-ch)
}
