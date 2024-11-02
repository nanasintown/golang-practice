//go:build ignore
// +build ignore

package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// j will be the job from <jobs>. more can be either true or false, if all jobs are received
			// and <jobs> is closed, more will be false
			if more {
				fmt.Println("Received job", j)
			} else {
				fmt.Println("Received all jobs")
				done <- true
				return
			}
		}
	}()
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent jobs", j)
	}
	// closing indicates that no more value can be sent to the channel, avoid deadlock,
	// especially when sending data to channel by using loop
	close(jobs)
	fmt.Println("sent all jobs")

	<-done

	_, ok := <-jobs
	fmt.Println("Received more jobs: ", ok)

	/*Range over channels*/
	queue := make(chan int, 2)
	queue <- 34
	queue <- 26
	close(queue)
	for ele := range queue {
		fmt.Println(ele)
	}

}
