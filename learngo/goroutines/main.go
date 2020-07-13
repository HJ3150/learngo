package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"nico", "flynn", "dal", "japanguy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	// fmt.Println("Waiting for message")
	for i := 0; i < len(people); i++ {
		fmt.Print("waiting for ", i)
		fmt.Println(<-c)
	}

	// go sexyCount("nico")
	// go sexyCount("flynn")
	// time.Sleep(time.Second * 5)
}

// func sexyCount(person string) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 10)
	c <- person + " is sexy"
}
