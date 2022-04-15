package main

import (
	"fmt"
	"time"
)

func main() {
	strs := []string{"one", "two", "three", "asdasd"}

	for _, s := range strs {
		go func() {
			time.Sleep(2 * time.Second)

			fmt.Printf("%s ", s)

		}()

	}
	time.Sleep(3 * time.Second)

}
