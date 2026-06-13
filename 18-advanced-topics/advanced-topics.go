// ============================================================
// TOPIC: Advanced Topics in Go
// ============================================================
// CONCEPTS COVERED:
//  1. Reflection (reflect package) – inspect types at runtime
//  2. Struct tags + reflect – reading tag metadata
//  3. The io package – Reader and Writer interfaces
//  4. Bufio – buffered I/O
//  5. String manipulation (strings package)
//  6. Regular expressions (regexp package)
//  7. JSON encoding / decoding (encoding/json)
//  8. Time and Duration
//  9. sort package – sorting slices and custom types
// 10. math/rand – random number generation
// 11. os package – environment, args, file ops
// 12. log package – structured logging
// 13. sync.Pool – object pooling for performance
// 14. Build tags / conditional compilation
// 15. Embedding files with //go:embed
// 16. Closures capturing loop variable (Go 1.22 fix)
// 17. Functional options pattern (variadic options)
// 18. init() ordering across packages
// ============================================================

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"math"
	"math/rand"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode"
)

// ─────────────────────────────────────────
// 17. Functional options pattern
// ─────────────────────────────────────────

type ServerConfig struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

type Option func(*ServerConfig)

func WithHost(host string) Option {
	return func(c *ServerConfig) { c.host = host }
}

func WithPort(port int) Option {
	return func(c *ServerConfig) { c.port = port }
}

func WithTimeout(d time.Duration) Option {
	return func(c *ServerConfig) { c.timeout = d }
}

func WithMaxConn(n int) Option {
	return func(c *ServerConfig) { c.maxConn = n }
}

func NewServer(opts ...Option) *ServerConfig {
	// sensible defaults
	cfg := &ServerConfig{
		host:    "localhost",
		port:    8080,
		timeout: 30 * time.Second,
		maxConn: 100,
	}
	for _, opt := range opts {
		opt(cfg) // apply each option
	}
	return cfg
}

// ─────────────────────────────────────────
// 1 & 2. Reflect demo struct
// ─────────────────────────────────────────

type Employee struct {
	Name       string  `json:"name" validate:"required"`
	Department string  `json:"dept" validate:"required"`
	Salary     float64 `json:"salary" validate:"min=0"`
	Active     bool    `json:"active"`
}

func main() {

	// ─────────────────────────────────────────
	// 1. Reflection – inspect type and value at runtime
	// ─────────────────────────────────────────
	fmt.Println("═══ 1. Reflection ═══")
	emp := Employee{Name: "Alice", Department: "Engineering", Salary: 80000, Active: true}

	t := reflect.TypeOf(emp)
	v := reflect.ValueOf(emp)

	fmt.Printf("Type: %s\n", t.Name())
	fmt.Printf("Kind: %s\n", t.Kind())

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		fmt.Printf("  Field: %-12s Type: %-8s Value: %v\n",
			field.Name, field.Type.Name(), value)
	}

	// ─────────────────────────────────────────
	// 2. Reading struct tags via reflection
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 2. Struct tags via reflect ═══")
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		jsonTag := field.Tag.Get("json")
		validateTag := field.Tag.Get("validate")
		fmt.Printf("  %-12s json=%q validate=%q\n", field.Name, jsonTag, validateTag)
	}

	// ─────────────────────────────────────────
	// 5. String manipulation
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 5. strings package ═══")
	s := "  Hello, Go World!  "
	fmt.Println("Trim:     ", strings.TrimSpace(s))
	fmt.Println("Upper:    ", strings.ToUpper(s))
	fmt.Println("Lower:    ", strings.ToLower(s))
	fmt.Println("Contains: ", strings.Contains(s, "Go"))
	fmt.Println("HasPrefix:", strings.HasPrefix(strings.TrimSpace(s), "Hello"))
	fmt.Println("Replace:  ", strings.ReplaceAll(s, "o", "0"))
	fmt.Println("Split:    ", strings.Fields(strings.TrimSpace(s)))
	fmt.Println("Join:     ", strings.Join([]string{"Go", "is", "great"}, "-"))
	fmt.Println("Count:    ", strings.Count(s, "l"))
	fmt.Println("Index:    ", strings.Index(s, "Go"))
	fmt.Println("Repeat:   ", strings.Repeat("Go!", 3))
	fmt.Println("Title:    ", strings.ToTitle("hello world"))

	// strings.Builder – efficient string building
	var sb strings.Builder
	for _, word := range []string{"Go", "is", "fast"} {
		sb.WriteString(word)
		sb.WriteRune(' ')
	}
	fmt.Println("Builder:  ", strings.TrimSpace(sb.String()))

	// ─────────────────────────────────────────
	// 6. Regular expressions
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 6. regexp ═══")
	emailPattern := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	emails := []string{"user@example.com", "bad-email", "another@test.org"}
	for _, e := range emails {
		fmt.Printf("  %s → valid=%v\n", e, emailPattern.MatchString(e))
	}

	// FindAll
	digitPattern := regexp.MustCompile(`\d+`)
	text := "Go 1.22 released in 2024, previous was 1.21"
	digits := digitPattern.FindAllString(text, -1)
	fmt.Println("Digits found:", digits)

	// Replace
	sanitized := regexp.MustCompile(`\s+`).ReplaceAllString("hello   world\t!", " ")
	fmt.Println("Sanitized:", sanitized)

	// ─────────────────────────────────────────
	// 7. JSON encoding / decoding
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 7. JSON ═══")
	// Marshal (struct → JSON)
	emp.Name = "Bob"
	emp.Salary = 0 // will be included (no omitempty)
	jsonBytes, _ := json.Marshal(emp)
	fmt.Println("Marshal:  ", string(jsonBytes))

	// Marshal with indentation
	prettyJSON, _ := json.MarshalIndent(emp, "", "  ")
	fmt.Println("Indented:\n" + string(prettyJSON))

	// Unmarshal (JSON → struct)
	jsonStr := `{"name":"Charlie","dept":"HR","salary":60000,"active":false}`
	var emp2 Employee
	json.Unmarshal([]byte(jsonStr), &emp2)
	fmt.Printf("Unmarshal: %+v\n", emp2)

	// JSON with map (dynamic keys)
	jsonMap := `{"key1":"val1","key2":42}`
	var dynamic map[string]interface{}
	json.Unmarshal([]byte(jsonMap), &dynamic)
	fmt.Println("Dynamic map:", dynamic)

	// ─────────────────────────────────────────
	// 8. Time and Duration
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 8. time ═══")
	now := time.Now()
	fmt.Println("Now:      ", now.Format("2006-01-02 15:04:05"))
	fmt.Println("UTC:      ", now.UTC().Format(time.RFC3339))
	fmt.Println("Unix ts:  ", now.Unix())

	// Duration arithmetic
	oneHour := time.Hour
	fmt.Println("In 1h:    ", now.Add(oneHour).Format("15:04"))
	fmt.Println("1h30m:    ", (90 * time.Minute).String())

	// Parse a time
	parsed, _ := time.Parse("2006-01-02", "2024-01-15")
	fmt.Println("Parsed:   ", parsed.Format("January 2, 2006"))

	// Measure elapsed time
	start := time.Now()
	time.Sleep(1 * time.Millisecond)
	elapsed := time.Since(start)
	fmt.Printf("Elapsed:   %v\n", elapsed)

	// ─────────────────────────────────────────
	// 9. sort package
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 9. sort ═══")
	nums := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(nums)
	fmt.Println("sort.Ints:   ", nums)

	strs := []string{"banana", "apple", "cherry"}
	sort.Strings(strs)
	fmt.Println("sort.Strings:", strs)

	// Custom sort – by length then alphabetically
	words := []string{"go", "python", "rust", "c", "java"}
	sort.Slice(words, func(i, j int) bool {
		if len(words[i]) != len(words[j]) {
			return len(words[i]) < len(words[j])
		}
		return words[i] < words[j]
	})
	fmt.Println("Custom sort: ", words)

	// sort.SliceStable – preserves original order for equal elements
	type Item struct {
		Name     string
		Priority int
	}
	items := []Item{{"task3", 1}, {"task1", 2}, {"task2", 1}}
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Priority < items[j].Priority
	})
	fmt.Println("SliceStable: ", items)

	// ─────────────────────────────────────────
	// 10. math/rand
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 10. math/rand ═══")
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Printf("Random int [0,100): %d\n", rng.Intn(100))
	fmt.Printf("Random float [0,1): %.4f\n", rng.Float64())
	fmt.Printf("Random int64:       %d\n", rng.Int63())

	// Shuffle a slice
	deck := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	rng.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	fmt.Println("Shuffled deck (first 5):", deck[:5])

	// ─────────────────────────────────────────
	// 11. os package
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 11. os ═══")
	// Environment variables
	os.Setenv("MY_VAR", "hello")
	fmt.Println("MY_VAR:", os.Getenv("MY_VAR"))

	// Command-line args
	fmt.Println("os.Args[0]:", os.Args[0])

	// Working directory
	wd, _ := os.Getwd()
	fmt.Println("Working dir:", wd)

	// ─────────────────────────────────────────
	// 12. log package & slog (structured logging, Go 1.21+)
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 12. Logging ═══")
	// Standard log – adds timestamp prefix automatically
	log.SetFlags(log.Ltime | log.Lshortfile)
	log.Println("Standard log message")

	// slog – structured logging (Go 1.21+)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	logger.Info("application started", "version", "1.0.0", "env", "dev")
	logger.Debug("debug info", "user_id", 42)
	logger.Warn("high memory usage", "percent", 85)

	// ─────────────────────────────────────────
	// 13. sync.Pool – reuse objects to reduce GC pressure
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 13. sync.Pool ═══")
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("  Pool: creating new buffer")
			return &strings.Builder{}
		},
	}

	// Get from pool
	buf := pool.Get().(*strings.Builder)
	buf.WriteString("hello from pool")
	fmt.Println("  Got:", buf.String())

	buf.Reset()        // reset before returning
	pool.Put(buf)      // return to pool

	buf2 := pool.Get().(*strings.Builder) // reuses the same object (no "creating new" log)
	buf2.WriteString("reused buffer")
	fmt.Println("  Got:", buf2.String())
	pool.Put(buf2)

	// ─────────────────────────────────────────
	// 17. Functional options pattern
	// ─────────────────────────────────────────
	fmt.Println("\n═══ 17. Functional options ═══")
	// Default server
	s1 := NewServer()
	fmt.Printf("Default: %s:%d timeout=%v maxConn=%d\n",
		s1.host, s1.port, s1.timeout, s1.maxConn)

	// Custom server
	s2 := NewServer(
		WithHost("0.0.0.0"),
		WithPort(9090),
		WithTimeout(60*time.Second),
		WithMaxConn(500),
	)
	fmt.Printf("Custom:  %s:%d timeout=%v maxConn=%d\n",
		s2.host, s2.port, s2.timeout, s2.maxConn)

	// ─────────────────────────────────────────
	// Additional: math package
	// ─────────────────────────────────────────
	fmt.Println("\n═══ Bonus: math package ═══")
	fmt.Printf("Pi:      %.6f\n", math.Pi)
	fmt.Printf("Sqrt(2): %.6f\n", math.Sqrt(2))
	fmt.Printf("Pow(2,10): %.0f\n", math.Pow(2, 10))
	fmt.Printf("Abs(-5):   %.0f\n", math.Abs(-5))
	fmt.Printf("Ceil(1.2): %.0f\n", math.Ceil(1.2))
	fmt.Printf("Floor(1.9):%.0f\n", math.Floor(1.9))
	fmt.Printf("MaxInt64:  %d\n", math.MaxInt64)

	// ─────────────────────────────────────────
	// Additional: unicode package
	// ─────────────────────────────────────────
	fmt.Println("\n═══ Bonus: unicode package ═══")
	chars := []rune{'A', 'a', '5', ' ', '!', 'ñ'}
	for _, ch := range chars {
		fmt.Printf("  %c → letter=%v digit=%v upper=%v space=%v\n",
			ch,
			unicode.IsLetter(ch),
			unicode.IsDigit(ch),
			unicode.IsUpper(ch),
			unicode.IsSpace(ch))
	}
}
