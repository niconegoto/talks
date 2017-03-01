package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	users := []string{"太郎", "次郎", "三郎"}

	var wg sync.WaitGroup

	for _, u := range users {

		wg.Add(1)
		go func() {
			fmt.Println("Hello", u)
			wg.Done()
		}()
		time.Sleep(1 * time.Second)
	}

	wg.Wait()
}
