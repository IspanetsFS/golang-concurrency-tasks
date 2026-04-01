# Go Concurrency Tasks

A collection of practical Go concurrency exercises designed to help you master goroutines, channels, and synchronization patterns used in real-world production code.

> The tasks were generated with the help of AI. I used them to prepare for technical live-coding interviews and to practice Go concurrency patterns.

Each task comes with:
- **`source/main.go`** — the skeleton with a task description and `TODO` stubs to implement
- **`solution/main.go`** — a reference implementation

## How to use

1. Read the task description at the top of `source/main.go`
2. Implement the `TODO` functions
3. Run your solution: `go run source/main.go`
4. Compare with the reference: `go run solution/main.go`

## Tasks

| Folder | Description |
|--------|-------------|
| [bounded-pipeline](./bounded-pipeline) | Implement a **pausable generator** using the nil-channel pattern. The generator must apply backpressure — pausing when the downstream consumer is not ready — instead of buffering or dropping values. |
| [circuit-breaker](./circuit-breaker) | Implement the **Circuit Breaker** pattern to protect against cascading failures. After `threshold` consecutive errors the circuit opens and rejects calls immediately; it resets after a configurable timeout. |
| [fan-out-fan-in](./fan-out-fan-in) | Implement **fan-out / fan-in**: distribute work from a single input channel across N worker goroutines (each squares its input), then merge all results back into one output channel. |
| [first-win](./first-win) | Launch N goroutines that each perform a request. Return the result of the **first one to succeed** and cancel all others via context. If every goroutine fails, return the last error. |
| [graceful-shutdown](./graceful-shutdown) | Implement a service that runs N background workers. When the context is cancelled, all workers must **finish gracefully** and the main goroutine must wait for all of them before returning. |
| [nilchannel](./nilchannel) | Implement a `merge` function that combines two channels into one using the **nil-channel pattern** (no `sync.WaitGroup` allowed). When one channel closes, keep reading the other; close the output only when both are done. |
| [once](./once) | Implement a **thread-safe lazy-loading cache**. On the first access to a key the loader function is called; subsequent accesses return the cached value. Concurrent requests for the same key must invoke the loader exactly once. |
| [pipeline](./pipeline) | Build a **three-stage pipeline**: `generate → square → filter`. The pipeline must shut down cleanly on context cancellation with no goroutine leaks. |
| [race-condition](./race-condition) | A counter with a **data race** is provided. Fix it in three independent ways: using `sync.Mutex`, using `sync/atomic`, and using a channel-based actor. |
| [rate-limiter](./rate-limiter) | Implement a **token-bucket rate limiter** backed by a ticker. `Wait` blocks until a token is available or the context is cancelled; `Stop` cleans up the underlying ticker. |
| [retry-timeout](./retry-timeout) | Implement a `Retry` helper with **exponential backoff**. It retries a function up to `MaxAttempts` times, doubling the wait on each failure, and respects context cancellation throughout. |
| [semaphore](./semaphore) | Implement a **semaphore** using a buffered channel. `Acquire` blocks when all slots are taken and returns early if the context is cancelled; `Release` frees a slot. |
| [singleflight](./singleflight) | Implement the **singleflight** pattern: if multiple goroutines concurrently request the same key, only one real call is made and all callers receive the same result. Unlike a cache, results are not stored after the call completes. |
| [worker-pool](./worker-pool) | Implement a **worker pool** that processes N jobs with at most M concurrent goroutines. Results must be returned in the **same order** as the input jobs. |

## Concepts covered

- Goroutines and `sync.WaitGroup`
- Unbuffered and buffered channels
- `select` statement and the nil-channel pattern
- Context cancellation (`context.WithCancel`, `context.WithTimeout`)
- `sync.Mutex` and `sync.RWMutex`
- `sync/atomic` operations
- `sync.Once`
- Backpressure and flow control
- Fan-out / fan-in
- Pipeline pattern
- Worker pool pattern
- Rate limiting
- Circuit Breaker pattern
- Singleflight / request deduplication
- Graceful shutdown

## Requirements

Go 1.22+ (uses range-over-integer and other modern features).

```bash
go version  # should be >= 1.22
```
