// ============================================================
// TOPIC: Goroutines in Go
// ============================================================
// CONCEPTS COVERED:
//  1. What is a goroutine? go keyword
//  2. Basic goroutine launch
//  3. sync.WaitGroup – wait for goroutines to finish
//  4. Race condition demo (goroutines sharing data unsafely)
//  5. sync.Mutex – mutual exclusion for shared data
//  6. sync.RWMutex – read-write lock
//  7. sync.Once – run something exactly once
//  8. sync.Atomic – atomic operations on integers
//  9. GOMAXPROCS – controlling parallelism
// 10. Goroutine with anonymous function (closure)
// 11. Worker pool pattern
// 12. Fan-out / fan-in (see also: 14-channels)
// ============================================================

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// ─────────────────────────────────────────
// Shared counter for race condition demo
// ─────────────────────────────────────────
type SafeCounter struct {
	mu    sync.Mutex
	count int
}

func (sc *SafeCounter) Inc() {
	sc.mu.Lock()   // acquire exclusive lock
	sc.count++
	sc.mu.Unlock() // release lock
}

func (sc *SafeCounter) Value() int {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	return sc.count
}

func main() {

	// ─────────────────────────────────────────
	// 1 & 2. Basic goroutine with go keyword
	// A goroutine is a lightweight thread managed by the Go runtime
	// ─────────────────────────────────────────
	fmt.Println("═══ 1 & 2. Basic goroutine ═══")
	go sayHello("Goroutine 1")
	go sayHello("Goroutine 2")

	// Without synchronization, main may exit before goroutines finish
	// We use time.Sleep here only for demonstration – use WaitGroup in real code
	time.Sleep(50 * time.Millisecond)

	// ─────────────────────────────────────────
	// 3. sync.WaitGroup – proper way to wait for goroutines
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 3. sync.WaitGroup ═══")
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)              // increment counter BEFORE launching goroutine
		go func(id int) {
			defer wg.Done()    // decrement counter when goroutine finishes
			fmt.Printf("  Worker %d done\n", id)
		}(i) // pass i as argument to avoid closure capture bug
	}

	wg.Wait() // block until counter reaches 0
	fmt.Println("All workers finished")

	// ─────────────────────────────────────────
	// 4. Race condition (UNSAFE – shown for education)
	// Run with: go run -race goroutines.go to detect races
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. Race condition (unsafe counter) ═══")
	unsafeCount := 0
	var wg2 sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			unsafeCount++ // DATA RACE: concurrent read+write
		}()
	}
	wg2.Wait()
	// Result is unpredictable – may not be 100!
	fmt.Printf("Unsafe count (may be wrong): %d\n", unsafeCount)

	// ─────────────────────────────────────────
	// 5. sync.Mutex – safe concurrent access
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. sync.Mutex (safe counter) ═══")
	sc := &SafeCounter{}
	var wg3 sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			sc.Inc()
		}()
	}
	wg3.Wait()
	fmt.Printf("Safe count (always 1000): %d\n", sc.Value())

	// ─────────────────────────────────────────
	// 6. sync.RWMutex – allows multiple concurrent readers
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. sync.RWMutex ═══")
	var rwmu sync.RWMutex
	data := map[string]string{"key": "initial"}

	// Many goroutines can read simultaneously
	var rwWg sync.WaitGroup
	for i := 0; i < 3; i++ {
		rwWg.Add(1)
		go func(id int) {
			defer rwWg.Done()
			rwmu.RLock()  // read lock (shared)
			val := data["key"]
			rwmu.RUnlock()
			fmt.Printf("  Reader %d got: %q\n", id, val)
		}(i)
	}
	// One writer at a time
	rwWg.Add(1)
	go func() {
		defer rwWg.Done()
		rwmu.Lock()   // write lock (exclusive)
		data["key"] = "updated"
		rwmu.Unlock()
		fmt.Println("  Writer updated the value")
	}()
	rwWg.Wait()

	// ─────────────────────────────────────────
	// 7. sync.Once – initialization runs exactly once
	// Useful for lazy singletons
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. sync.Once ═══")
	var once sync.Once
	var onceWg sync.WaitGroup
	for i := 0; i < 5; i++ {
		onceWg.Add(1)
		go func(id int) {
			defer onceWg.Done()
			once.Do(func() {
				fmt.Printf("  Initializing… (called by goroutine %d)\n", id)
			})
		}(i)
	}
	onceWg.Wait()
	fmt.Println("  sync.Once: init ran exactly once despite 5 goroutines")

	// ─────────────────────────────────────────
	// 8. sync/atomic – lock-free atomic operations
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. sync/atomic ═══")
	var atomicCounter int64
	var atomicWg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		atomicWg.Add(1)
		go func() {
			defer atomicWg.Done()
			atomic.AddInt64(&atomicCounter, 1) // atomic: no lock needed for simple ops
		}()
	}
	atomicWg.Wait()
	fmt.Printf("Atomic counter (always 1000): %d\n", atomic.LoadInt64(&atomicCounter))

	// ─────────────────────────────────────────
	// 10. Goroutine with closure – capture loop variable correctly
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. Goroutine with closure (correct capture) ═══")
	var closureWg sync.WaitGroup
	for i := 0; i < 3; i++ {
		closureWg.Add(1)
		i := i // shadow i – creates a new variable per iteration (Go 1.21 fixes this)
		go func() {
			defer closureWg.Done()
			fmt.Printf("  Closure goroutine: i = %d\n", i)
		}()
	}
	closureWg.Wait()

	// ─────────────────────────────────────────
	// 11. Worker pool pattern
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. Worker pool ═══")
	workerPool(3, 9) // 3 workers, 9 jobs
}

// ─────────────────────────────────────────
// Helper functions
// ─────────────────────────────────────────

func sayHello(name string) {
	fmt.Printf("  Hello from %s!\n", name)
}

// 11. Worker pool using channels and WaitGroup
func workerPool(numWorkers, numJobs int) {
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Start workers
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for job := range jobs { // worker pulls from jobs channel
				result := job * job // "work": square the number
				fmt.Printf("  Worker %d processed job %d → %d\n", workerID, job, result)
				results <- result
			}
		}(w)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // signal workers: no more jobs

	// Close results after all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	total := 0
	for r := range results {
		total += r
	}
	fmt.Printf("Worker pool total: %d\n", total)
}
