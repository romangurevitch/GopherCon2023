# Working with Goroutines

## Table of contents

### Introduction

During the demo I will ask you 3 questions:

1. Did we get the expected result?
2. Do we have any race conditions?
3. Did we handle all the errors and shutdowns gracefully?

### Counter

* [Basic counter](../../docs/goroutine/counter/basic.md)
* [Mutex counter](../../docs/goroutine/counter/mutex.md)
* [RWMutex counter](../../docs/goroutine/counter/rwmutex.md)
* [Atomic counter](../../docs/goroutine/counter/atomic.md)

### Race detection flag

* [Race detection](../../docs/goroutine/race/race.md)

### Concurrency building blocks

* [WaitGroup](../concurrency/sync/waitgroup/README.md)
* [Locks](../concurrency/sync/locks/README.md)
* [Atomic](../concurrency/sync/atomic/README.md)
* [Channels](../concurrency/channel/README.md)
* [Signals](../concurrency/signal/README.md)
* [Context](../concurrency/context/README.md)

### Examples

* [1. Let's start with a basic example](../../docs/goroutine/example_1.md)
* [2. Let's try and fix the issues](../../docs/goroutine/example_2.md)
* [3. Let's make a small change :)](../../docs/goroutine/example_3.md)
* [4. Finally, something works as expected](../../docs/goroutine/example_4.md)
* [5. Non-stopping go routines](../../docs/goroutine/example_5.md)
* [6. Let's handle shutdown gracefully?](../../docs/goroutine/example_6.md)
* [7. Let's handle shutdown gracefully, for real this time!](../../docs/goroutine/example_7.md)
* [8. Adding context](../../docs/goroutine/example_8.md)
* [9. Notify context on signals](../../docs/goroutine/example_9.md)
* [8. Bonus: let's make a tiny change](../../docs/goroutine/example_9_bonus.md)

### Questions?

* [Question?](../../docs/goroutine/questions.md)