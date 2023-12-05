package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// concurrency control
	// useWaitGroup()
	// useChannelSelect()
	// useContext()
	useUseWorker()
}

func useWaitGroup() {
	var wg sync.WaitGroup
	wg.Add(2) // how much task to be wait
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Job 1 done.")
		wg.Done()
	}()
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("Job 2 done.")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("All Done.")
}

func useChannelSelect() {
	stop := make(chan bool)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("got stop channel")
				return
			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	stop <- true
	time.Sleep(5 * time.Second)
}

func useContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("got the stop channel")
				return
			default:
				fmt.Println("working")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("stop the gorutine")
	cancel()
	time.Sleep(5 * time.Second)
}

func useUseWorker() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx, "job1")
	go worker(ctx, "job2")
	go worker(ctx, "job3")

	time.Sleep(5 * time.Second)
	fmt.Println("stop the jobs")
	cancel()
	time.Sleep(5 * time.Second)
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "got the stop channel")
			return
		default:
			fmt.Println(name, "working")
			time.Sleep(1 * time.Second)
		}
	}
}
