A sync.WaitGroup is a synchronization primitive from the sync package in Go that can be used to wait for a collection of goroutines
to finish executing. It is particularly useful in scenarios where you have multiple concurrent operations
and you need to wait for all of them to complete before proceeding. The sync.WaitGroup provides a simple way to orchestrate this synchronization.

How to Use sync.WaitGroup
A WaitGroup uses a counter to keep track of how many goroutines are currently running. You manipulate this counter with three main methods:

Add(int): Increments the WaitGroup counter by the specified value. You typically call this before starting a goroutine, to indicate that another goroutine has begun.
Done(): Decrements the WaitGroup counter by one, indicating that a goroutine has finished. It's equivalent to calling Add(-1).
Wait(): Blocks until the WaitGroup counter returns to zero, i.e., all goroutines have finished.

===================================================================================================================================================

The RWMutex in Go, part of the sync package, is a reader/writer mutual exclusion lock.
The name "RWMutex" stands for Read/Write Mutex. This type of mutex allows multiple readers to hold the lock simultaneously or a single writer
to hold the lock exclusively. This distinction makes RWMutex particularly useful in situations where data structures are read frequently
but written to infrequently.

How RWMutex Works
Read Lock (RLock): When a goroutine locks an RWMutex for reading (using RLock()), other goroutines can also read lock (RLock()) it simultaneously,
but no goroutine can write lock (Lock()) it until all the read locks are released.

Write Lock (Lock): When a goroutine locks an RWMutex for writing (using Lock()), it gets exclusive access.
No other goroutine can either read lock (RLock()) or write lock (Lock()) it until the writer releases the lock.

This behavior optimizes performance in read-heavy scenarios by allowing concurrent reads,
which can significantly improve efficiency compared to using a standard Mutex.