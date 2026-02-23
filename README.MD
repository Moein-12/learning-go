# Personal Go Learning Projects

A collection of systems programming projects built from scratch to learn Go's concurrency model, standard library, and tooling mindset. Each project is implemented independently, then iteratively improved—no AI-generated code, just deliberate practice.

---

## Philosophy

### Stage: Iteration 1 - Project 3

I'm coming from Python/Rust and system design, preferring small, algorithmic developer tools over large applications. These projects are my hands-on way to internalize Go idioms, concurrency patterns, and the "standard library first" mentality.

**The rules:**

- Must be done in this order
- After the last project is done, I Must go back and improve my codes or introduce a new code from the beginning
- This projects ends when i get all 7 in 5 commit or less each and i have learned all that's mentioned

Notes:

- I did two first projects off github so i don't count the first iteration at all
- The projects are designed by DeepSeek based on my background in programming

---

## Projects

Each project targets specific Go concepts and builds on previous learnings. Implemented in order, then revisited for improvements.

### 1. `wc` Clone — Word Count CLI

*Unix-style word count tool for files and stdin*

**Concepts:**

- CLI argument parsing with `flag`
- File I/O (`os.Open`, `bufio.Scanner`)
- Unicode handling with `unicode/utf8`
- Idiomatic error returns
- Basic `Makefile` and build workflow

**Extensions:** Multiple file support, total aggregation

---

### 2. Concurrent Prime Sieve

*Generate primes up to N using goroutines + channels*

**Concepts:**

- Goroutines and channel communication
- Pipeline pattern (generator → filter → collector)
- `sync.WaitGroup` for coordination
- Benchmarking with `go test -bench`
- Race condition detection with `-race`

**Extensions:** HTTP endpoint for streaming primes

---

### 3. HTTP Load Tester

*Concurrent HTTP benchmarking tool with latency stats*

**Concepts:**

- `net/http` client configuration and timeouts
- Worker pool pattern with channels
- `sync.Mutex` for shared state protection
- Duration calculations and percentile stats
- Graceful shutdown (signal handling)

**Extensions:** Custom headers, POST data, JSON output

---

### 4. In-Memory Cache with TTL & Persistence

*Thread-safe key-value store with expiration and disk backup*

**Concepts:**

- `map` with `sync.RWMutex` for concurrency safety
- Background cleanup with `time.Ticker`
- `encoding/gob`/`encoding/json` for persistence
- Struct tags for serialization
- Integration tests with temporary files

**Extensions:** LRU eviction, memcached-like protocol

---

### 5. TCP Chat Server

*Multi-client chat room with broadcast messaging*

**Concepts:**

- `net.Listen` and per-connection goroutines
- Line-based protocols with `bufio`
- Hub pattern for broadcast coordination
- Connection lifecycle management
- `context` for timeouts and cancellation

**Extensions:** Private messages, rooms, HTTP stats endpoint

---

### 6. Custom Static Analysis Linter

*AST-based tool enforcing custom Go style rules*

**Concepts:**

- `go/parser`, `go/ast`, `go/token` traversal
- Type information with `go/types`
- `go/analysis` framework integration
- Testing analysis with testdata
- `go vet`-compatible tooling

**Extensions:** Context cancellation detection, slice efficiency checks

---

### 7. Simplified Raft KV Store

*Distributed key-value store with consensus simulation*

**Concepts:**

- RPC communication (`net/rpc` or HTTP)
- Timer-based leader election
- `sync.Cond` and channel coordination
- State machine design
- Deterministic testing for distributed systems

**Extensions:** Multi-process replication

---

### Bonus: Data Structures from Scratch

*Algorithmic deep dives with Go-specific considerations*

- Concurrent hash map (sharded locking)
- Bloom filter
- Trie with prefix search
- Consistent hashing ring

---

## My definition of Done

If it works correctly, it's done. I'm not holding myself back so that I do it the way described, but I'll go back to them eventually.
