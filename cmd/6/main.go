// Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker представляет собой структуру для работы с горутинами.
type Worker struct {
	wg *sync.WaitGroup
}

// CheckChannelClose проверяет закрылся ли канал, пытаясь вычитать из него данные.
// Если канал закрыт, функция завершается.
func (w *Worker) CheckChannelClose(ch <-chan int) {
	defer w.wg.Done()
	for {
		v, ok := <-ch
		if !ok {
			fmt.Println("Worker: CheckChannelClose finished")
			return
		}
		fmt.Printf("Worker: got value from channel: %d\n", v)
	}
}

// ChannelCloseRange завершится, когда канал закроется.
func (w *Worker) ChannelCloseRange(ch <-chan int) {
	defer w.wg.Done()
	for v := range ch {
		fmt.Printf("Worker: got value from channel: %d\n", v)
	}
	fmt.Println("Worker: ChannelCloseRange finished")
}

// ChannelCloseWithCancel использует отдельный канал для завершения горутины.
func (w *Worker) ChannelCloseWithCancel(ch <-chan int, stop <-chan struct{}) {
	defer w.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("Worker: got value from channel: %d\n", v)
		case <-stop:
			fmt.Println("Worker: ChannelCloseWithCancel finished")
			return
		}
	}
}

// ChannelCloseWithContext завершается по сигналу из контекста.
func (w *Worker) ChannelCloseWithContext(ctx context.Context, ch <-chan int) {
	defer w.wg.Done()
	for {
		select {
		case v := <-ch:
			fmt.Printf("Worker: got value from channel: %d\n", v)
		case <-ctx.Done():
			fmt.Println("Worker: ChannelCloseWithContext finished")
			return
		}
	}
}

// ChannelCloseWithContextTimeout завершается по сигналу из контекста или по таймауту.
func (w *Worker) ChannelCloseWithContextTimeout(ctx context.Context, ch <-chan int) {
	defer w.wg.Done()
	for {
		select {
		case <-time.After(10 * time.Second): // долгая операция
			fmt.Printf("Worker: got value from channel: %d\n", <-ch)
		case <-ctx.Done():
			fmt.Println("Worker: ChannelCloseWithContextTimeout finished")
			return
		}
	}
}

// Producer записывает случайные числа в канал с определенной периодичностью.
// Завершается по сигналу из контекста.
func (w *Worker) Producer(ctx context.Context, ch chan<- int) {
	defer w.wg.Done()
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			v := rand.Int()
			ch <- v
			fmt.Printf("Producer: sent value to channel: %d\n", v)
		case <-ctx.Done():
			fmt.Println("Producer: exiting")
			return
		}
	}
}

func main() {
	fmt.Println("Context timeout exceeded")
	ctx, _ := context.WithCancel(context.Background())
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Важно вызвать cancel, чтобы избежать утечки контекста

	ch := make(chan int, 1)
	worker := &Worker{wg: &sync.WaitGroup{}}
	worker.wg.Add(1)
	go worker.Producer(timeoutCtx, ch)
	go worker.ChannelCloseWithContextTimeout(timeoutCtx, ch)
	worker.wg.Wait()

	fmt.Println("\nContext cancel")
	ctx, cancel = context.WithCancel(context.Background())
	ch = make(chan int, 1)
	worker.wg = &sync.WaitGroup{}
	worker.wg.Add(2)
	go worker.Producer(ctx, ch)
	go worker.ChannelCloseWithContext(ctx, ch)
	time.Sleep(3 * time.Second)
	cancel()
	close(ch)
	worker.wg.Wait()

	fmt.Println("\nStop channel")
	ch = make(chan int, 1)
	stop := make(chan struct{})
	ctx, cancel = context.WithCancel(context.Background())
	worker.wg = &sync.WaitGroup{}
	worker.wg.Add(2)
	go worker.Producer(ctx, ch)
	go worker.ChannelCloseWithCancel(ch, stop)
	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	cancel()
	close(ch)
	worker.wg.Wait()

	fmt.Println("\nRanging channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())
	worker.wg = &sync.WaitGroup{}
	worker.wg.Add(2)
	go worker.Producer(ctx, ch)
	go worker.ChannelCloseRange(ch)
	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	worker.wg.Wait()

	fmt.Println("\nClosing channel")
	ch = make(chan int, 1)
	ctx, cancel = context.WithCancel(context.Background())
	worker.wg = &sync.WaitGroup{}
	worker.wg.Add(2)
	go worker.Producer(ctx, ch)
	go worker.CheckChannelClose(ch)
	time.Sleep(3 * time.Second)
	close(ch)
	cancel()
	worker.wg.Wait()
}
