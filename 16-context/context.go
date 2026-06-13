// ============================================================
// TOPIC: Context in Go
// ============================================================
// CONCEPTS COVERED:
//  1. What is context.Context?
//  2. context.Background() – root context (never cancelled)
//  3. context.TODO() – placeholder, don't know context yet
//  4. context.WithCancel() – manual cancellation
//  5. context.WithTimeout() – cancel after duration
//  6. context.WithDeadline() – cancel at absolute time
//  7. context.WithValue() – carry request-scoped values
//  8. Propagating context through function calls
//  9. Checking ctx.Done() in a goroutine
// 10. ctx.Err() – reason for cancellation
// 11. Context in HTTP handlers (real-world use)
// 12. Best practices: always pass ctx as first argument
// ============================================================

package main

import (
	"context"
	"fmt"
	"time"
)

// ─────────────────────────────────────────
// 7. Custom context key type (prevent collisions)
// Always use unexported custom types for context keys
// ─────────────────────────────────────────
type contextKey string

const (
	userIDKey    contextKey = "userID"
	requestIDKey contextKey = "requestID"
)

func main() {

	// ─────────────────────────────────────────
	// 2. context.Background() – top-level root
	// ─────────────────────────────────────────
	fmt.Println("═══ 2. context.Background() ═══")
	bg := context.Background()
	fmt.Printf("Background: %v\n", bg)

	// ─────────────────────────────────────────
	// 3. context.TODO() – use when unsure which context to use
	// ─────────────────────────────────────────
	todo := context.TODO()
	fmt.Printf("TODO: %v\n", todo)

	// ─────────────────────────────────────────
	// 4. context.WithCancel() – manual cancellation
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 4. WithCancel ═══")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				// 10. ctx.Err() tells WHY context was cancelled
				fmt.Println("  Goroutine: context cancelled:", ctx.Err())
				return
			default:
				fmt.Println("  Goroutine: working…")
				time.Sleep(30 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(80 * time.Millisecond)
	cancel() // manually cancel – signals all children
	time.Sleep(20 * time.Millisecond)
	fmt.Println("  Main: called cancel()")

	// ─────────────────────────────────────────
	// 5. context.WithTimeout()
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. WithTimeout ═══")
	ctxTimeout, cancelTimeout := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancelTimeout() // always defer cancel to free resources

	result := slowOperation(ctxTimeout, 200*time.Millisecond) // slower than timeout
	fmt.Println("  Result:", result)

	result2 := slowOperation(context.Background(), 50*time.Millisecond) // no timeout
	fmt.Println("  Result2:", result2)

	// ─────────────────────────────────────────
	// 6. context.WithDeadline() – absolute point in time
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. WithDeadline ═══")
	deadline := time.Now().Add(150 * time.Millisecond)
	ctxDeadline, cancelDeadline := context.WithDeadline(context.Background(), deadline)
	defer cancelDeadline()

	select {
	case <-time.After(50 * time.Millisecond):
		fmt.Println("  Finished before deadline")
	case <-ctxDeadline.Done():
		fmt.Println("  Deadline exceeded:", ctxDeadline.Err())
	}

	// ─────────────────────────────────────────
	// 7. context.WithValue() – pass request-scoped data
	// Use sparingly: only for cross-cutting concerns (request IDs, auth tokens)
	// NOT for passing optional function parameters
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. WithValue ═══")
	ctxVal := context.WithValue(context.Background(), userIDKey, 42)
	ctxVal = context.WithValue(ctxVal, requestIDKey, "req-abc-123")

	processRequest(ctxVal)

	// ─────────────────────────────────────────
	// 8 & 9. Propagating context through call chain
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8 & 9. Context propagation ═══")
	rootCtx, rootCancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer rootCancel()
	orchestrate(rootCtx)

	// ─────────────────────────────────────────
	// 10. ctx.Err() – check why context ended
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. ctx.Err() ═══")
	ctxE, cancelE := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelE()
	time.Sleep(20 * time.Millisecond) // wait for timeout

	if ctxE.Err() == context.DeadlineExceeded {
		fmt.Println("  Timed out (DeadlineExceeded)")
	} else if ctxE.Err() == context.Canceled {
		fmt.Println("  Manually cancelled")
	} else {
		fmt.Println("  No error")
	}

	// ─────────────────────────────────────────
	// 12. Best practices summary
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Best Practices ═══")
	fmt.Println("  ✓ Pass ctx as the FIRST argument to every function that needs it")
	fmt.Println("  ✓ Never store context in structs (pass explicitly)")
	fmt.Println("  ✓ Always defer cancel() to release resources")
	fmt.Println("  ✓ Check ctx.Done() in long-running loops")
	fmt.Println("  ✓ Use typed keys for WithValue to avoid key collisions")
	fmt.Println("  ✗ Don't use context.Background() deep in call chains")
	fmt.Println("  ✗ Don't pass nil context; use TODO() as placeholder")
}

// ─────────────────────────────────────────
// Helper functions
// ─────────────────────────────────────────

// 5. Function that respects context cancellation
func slowOperation(ctx context.Context, duration time.Duration) string {
	select {
	case <-time.After(duration):
		return "success"
	case <-ctx.Done():
		return fmt.Sprintf("cancelled: %v", ctx.Err())
	}
}

// 7. Extract values from context
func processRequest(ctx context.Context) {
	userID, ok := ctx.Value(userIDKey).(int)
	if !ok {
		fmt.Println("  No userID in context")
		return
	}
	requestID, _ := ctx.Value(requestIDKey).(string)
	fmt.Printf("  Processing request %s for user %d\n", requestID, userID)
}

// 8. Passes context down the call chain
func orchestrate(ctx context.Context) {
	// Create a child context with additional timeout
	childCtx, cancel := context.WithTimeout(ctx, 300*time.Millisecond)
	defer cancel()

	done := make(chan string, 1)
	go func() {
		// Simulate work
		time.Sleep(50 * time.Millisecond)
		done <- "sub-task complete"
	}()

	select {
	case result := <-done:
		fmt.Println(" ", result)
	case <-childCtx.Done():
		fmt.Println("  Child context cancelled:", childCtx.Err())
	}
}
