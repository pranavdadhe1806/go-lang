// ============================================================
// TOPIC: Channels in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Unbuffered channel – synchronous send/receive
//  2. Buffered channel – async up to buffer capacity
//  3. Channel direction – send-only (<-chan) and receive-only (chan<-)
//  4. Closing a channel
//  5. Range over a channel
//  6. select statement – multiplex channels
//  7. select with default – non-blocking operations
//  8. select with timeout (time.After)
//  9. Done channel pattern (graceful shutdown signal)
// 10. Fan-out – distribute work across goroutines
// 11. Fan-in (merge) – combine multiple channels into one
// 12. Pipeline pattern
// 13. Semaphore pattern (limit concurrency)
// 14. Channel as a mutex (token passing)
// ============================================================

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// ─────────────────────────────────────────
	// 1. Unbuffered channel
	// Send blocks until someone receives; receive blocks until someone sends
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. Unbuffered channel ═══")
	ch := make(chan string) // unbuffered

	go func() {
		ch <- "Hello from goroutine!" // blocks until main receives
	}()

	msg := <-ch // blocks until goroutine sends
	fmt.Println("Received:", msg)

	// ─────────────────────────────────────────
	// 2. Buffered channel
	// Send is non-blocking when buffer has space
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. Buffered channel ═══")
	bch := make(chan int, 3) // buffer of 3

	bch <- 10 // does not block (buffer has space)
	bch <- 20
	bch <- 30
	// bch <- 40  // would block: buffer full

	fmt.Println("Buffered recv 1:", <-bch)
	fmt.Println("Buffered recv 2:", <-bch)
	fmt.Println("Buffered recv 3:", <-bch)

	// ─────────────────────────────────────────
	// 3. Channel directions – restrict channel usage
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. Channel direction ═══")
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	go pingPong(ping, pong)
	ping <- "ping"
	result := <-pong
	fmt.Println("Got back:", result)

	// ─────────────────────────────────────────
	// 4. Closing a channel
	// Sender closes; receivers get zero value + ok=false
	// Never close from the receiver side!
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Closing a channel ═══")
	numCh := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		numCh <- i
	}
	close(numCh) // sender closes

	for {
		v, ok := <-numCh
		if !ok {
			fmt.Println("Channel closed, done receiving")
			break
		}
		fmt.Println("Received:", v)
	}

	// ─────────────────────────────────────────
	// 5. Range over a channel (auto-handles close)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. Range over channel ═══")
	letters := make(chan string, 3)
	go func() {
		for _, l := range []string{"a", "b", "c"} {
			letters <- l
		}
		close(letters) // MUST close; otherwise range blocks forever
	}()
	for l := range letters {
		fmt.Println("Letter:", l)
	}

	// ─────────────────────────────────────────
	// 6. select – multiplex channels (waits on whichever is ready)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. select ═══")
	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func() { time.Sleep(10 * time.Millisecond); c1 <- "from c1" }()
	go func() { c2 <- "from c2" }()

	time.Sleep(20 * time.Millisecond) // let goroutines run
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Selected:", msg1)
		case msg2 := <-c2:
			fmt.Println("Selected:", msg2)
		}
	}

	// ─────────────────────────────────────────
	// 7. select with default – non-blocking channel ops
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. select with default ═══")
	nonBlock := make(chan int)
	select {
	case v := <-nonBlock:
		fmt.Println("Received:", v)
	default:
		fmt.Println("No message ready (non-blocking)")
	}

	// ─────────────────────────────────────────
	// 8. select with timeout using time.After
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. Timeout with select ═══")
	slowCh := make(chan string)
	go func() {
		time.Sleep(200 * time.Millisecond)
		slowCh <- "finally done"
	}()

	select {
	case result := <-slowCh:
		fmt.Println("Got:", result)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("Timed out! Channel was too slow.")
	}

	// ─────────────────────────────────────────
	// 9. Done channel – signal cancellation
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. Done channel ═══")
	done := make(chan struct{}) // empty struct: zero memory

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("  Worker: received done signal, stopping")
				return
			default:
				// do work...
			}
		}
	}()

	time.Sleep(10 * time.Millisecond)
	close(done) // broadcast to ALL receivers at once
	time.Sleep(10 * time.Millisecond)
	fmt.Println("  Main: sent done signal")

	// ─────────────────────────────────────────
	// 12. Pipeline pattern
	// Data flows through a series of stages connected by channels
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Pipeline pattern ═══")
	// Stage 1: generate numbers
	naturals := generate(1, 2, 3, 4, 5)
	// Stage 2: square them
	squares := squarer(naturals)
	// Stage 3: print them
	for s := range squares {
		fmt.Printf("  %d\n", s)
	}

	// ─────────────────────────────────────────
	// 11. Fan-in (merge multiple channels)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Fan-in ═══")
	ch1 := producer("A", 3)
	ch2 := producer("B", 3)
	merged := fanIn(ch1, ch2)
	for i := 0; i < 6; i++ {
		fmt.Println(" ", <-merged)
	}

	// ─────────────────────────────────────────
	// 13. Semaphore – limit concurrent goroutines
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. Semaphore (max 3 concurrent) ═══")
	sem := make(chan struct{}, 3) // allows max 3 concurrent
	var semWg sync.WaitGroup
	for i := 1; i <= 8; i++ {
		semWg.Add(1)
		go func(id int) {
			defer semWg.Done()
			sem <- struct{}{} // acquire slot
			fmt.Printf("  Task %d running\n", id)
			time.Sleep(5 * time.Millisecond)
			<-sem              // release slot
		}(i)
	}
	semWg.Wait()
	fmt.Println("All semaphore tasks done")
}

// ─────────────────────────────────────────
// 3. Directional channel functions
// ─────────────────────────────────────────
func pingPong(in <-chan string, out chan<- string) {
	msg := <-in           // receive-only
	out <- msg + " → pong" // send-only
}

// ─────────────────────────────────────────
// 12. Pipeline stage functions
// ─────────────────────────────────────────

// generate sends values into a channel and closes it
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// squarer reads from in, squares each value, sends to out
func squarer(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}

// ─────────────────────────────────────────
// 11. Fan-in helpers
// ─────────────────────────────────────────

func producer(name string, n int) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 1; i <= n; i++ {
			ch <- fmt.Sprintf("%s-%d", name, i)
		}
		close(ch)
	}()
	return ch
}

func fanIn(channels ...<-chan string) <-chan string {
	merged := make(chan string)
	var wg sync.WaitGroup
	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan string) {
			defer wg.Done()
			for v := range c {
				merged <- v
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		close(merged)
	}()
	return merged
}
