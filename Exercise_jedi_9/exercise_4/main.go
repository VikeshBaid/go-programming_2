package main

//race condition resolve using sync.Mutex
import (
	"fmt"
	"sync"
)

func main() {
  var wg sync.WaitGroup

  counter := 0
  gs := 10

	wg.Add(gs)

	var m sync.Mutex

  for i := 0; i < gs; i++ {
    go func ()  {
      m.Lock()
      v1 := counter
      fmt.Println("v1 counter:", v1)
      v1++
      counter = v1
      fmt.Println("v1 counter after inc: ", v1)
      m.Unlock()
      wg.Done()
    }()
  }

  wg.Wait()
	fmt.Println("count: ", counter)
}
