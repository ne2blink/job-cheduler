package job_test

import (
	"fmt"
	"sync"

	"github.com/ne2blink/job-scheduler"
)

func ExampleQueueScheduler() {
	s := job.NewQueueScheduler(1)

	var wg sync.WaitGroup
	wg.Add(3)
	s.Add(func() {
		defer wg.Done()
		fmt.Printf("hello")
	})
	s.Add(func() {
		defer wg.Done()
		fmt.Printf(" ")
	})
	s.Add(func() {
		defer wg.Done()
		fmt.Println("world")
	})

	wg.Wait()
	// Output: hello world
}
