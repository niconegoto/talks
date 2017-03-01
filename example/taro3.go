package main

import (
	"fmt"
	"sync"
)

func main() {
	users := []string{"太郎", "次郎", "三郎"}

	var wg sync.WaitGroup

	for _, u := range users {

		u := u
		wg.Add(1)
		go func() {
			fmt.Println("Hello", u)
			wg.Done()
		}()
	}

	wg.Wait()
}
