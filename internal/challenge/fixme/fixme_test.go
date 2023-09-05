package fixme

import (
	"context"
	"errors"
	"log/slog"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"sync/atomic"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"golang.org/x/sync/errgroup"
)

func TestBasicWriteDeadlock(t *testing.T) {
	exitAfter(time.Millisecond)

	ch := make(chan int)
	ch <- 42
	slog.Info("successfully sent on channel")
}

func TestBasicNilChannel(t *testing.T) {
	exitAfter(time.Millisecond)

	var ch chan int

	go func() {
		ch <- 1
		close(ch)
	}()

	for val := range ch {
		slog.Info("successfully received", "value", val)
	}
}

// nolint
func TestBasicClosedChannelWithoutOkCheck(t *testing.T) {
	exitAfter(time.Millisecond)
	ch := make(chan int)

	go func() {
		ch <- 42
		close(ch)
	}()

	for {
		select {
		case val := <-ch:
			slog.Info("received", "value", val)
		}
	}
}

func TestBasicUnlockingUnlockedLock(t *testing.T) {
	var mu sync.Mutex
	mu.Unlock()
}

func TestBasicWaitGroupNegativeCounter(t *testing.T) {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		wg.Done()
		wg.Done()
	}()

	wg.Wait()
}

// nolint
func TestBasicContextUsingPrimitivesAsKeys(t *testing.T) {
	type ctxKey string
	const key ctxKey = "myKey"
	ctx := context.WithValue(context.Background(), "myKey", "value1")

	if val, ok := ctx.Value(key).(string); !ok || val != "value1" {
		t.Fatalf("expected context to have 'value1' for 'myKey', got: %v", val)
	}
}

// nolint
func TestIntermediateUnbufferedNotifyChannel(t *testing.T) {
	exitAfter(100 * time.Millisecond)

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGINT)

	go func() {
		time.Sleep(10 * time.Millisecond)
		// Simulate sending a SIGINT to our own process
		if err := syscall.Kill(syscall.Getpid(), syscall.SIGINT); err != nil {
			require.NoError(t, err, "failed to send SIGINT")
		}
	}()

	select {
	case <-sigCh:
	}
}

func TestIntermediateDeadlock(t *testing.T) {
	exitAfter(100 * time.Millisecond)

	var mu sync.Mutex
	mu.Lock()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
	}()

	wg.Wait()
	slog.Error("success")
}

// nolint
func TestIntermediateWaitGroupByValue(t *testing.T) {
	exitAfter(100 * time.Millisecond)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func(wg sync.WaitGroup) {
		defer wg.Done()
	}(wg)

	wg.Wait()
}

// nolint
func TestIntermediateWaitGroupIncorrectAdd(t *testing.T) {
	wg := sync.WaitGroup{}
	finishedSuccessfully := false

	go func() {
		wg.Add(1)
		defer wg.Done()
		defer func() {
			finishedSuccessfully = true
		}()
	}()

	wg.Wait()
	require.True(t, finishedSuccessfully)
}

func TestIntermediateDefaultBusyLoop(t *testing.T) {
	ch := make(chan int)

	go func() {
		for i := 0; i < 3; i++ {
			ch <- 1
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()

	counter := 0
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				return
			}
			slog.Info("received", "value", val)
		default:
			counter++
			require.Less(t, counter, 100, "too many default calls")
		}
	}
}

func TestIntermediateMixingAtomicAndNonAtomicOperations(t *testing.T) {
	var count int32
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&count, 1)
		}()
	}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}

	wg.Wait()
	require.Equal(t, int32(2000), count, "Count was not updated atomically")
}

func TestIntermediateUnorderedReadFromChannels(t *testing.T) {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 2
	ch2 <- 3

	result := 5
	for i := 0; i < 2; i++ {
		select {
		case val := <-ch1:
			result *= val // result * 2
		case val := <-ch2:
			result += val // result + 3
		}
	}

	expected := 1
	require.Equal(t, expected, result)
}

func TestAdvancedWaitGroupWithoutDefer(t *testing.T) {
	exitAfter(100 * time.Millisecond)

	wg := sync.WaitGroup{}
	finishedSuccessfully := false

	finishedFunc := func() {
		finishedSuccessfully = true
		runtime.Goexit()
	}

	wg.Add(1)
	go func() {
		finishedFunc()
		wg.Done()
	}()

	wg.Wait()
	require.True(t, finishedSuccessfully)
}

// nolint
func TestAdvancedErrGroupWithoutWithContext(t *testing.T) {
	exitAfter(10 * time.Millisecond)
	expectedErr := errors.New("error")
	ctx := context.Background()
	group := errgroup.Group{}

	group.Go(func() error {
		return expectedErr
	})

	group.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	})

	if err := group.Wait(); err != nil {
		require.ErrorIs(t, err, expectedErr)
	}
}

// nolint
func TestAdvancedContextIgnoringCancellation(t *testing.T) {
	exitAfter(10 * time.Millisecond)
	_, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	inputCh := make(chan bool)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		select {
		// Waiting on input
		case <-inputCh:
		}
	}()

	wg.Wait()
}

func TestAdvancedMultipleProducersCloseChannel(t *testing.T) {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	producer := func() {
		defer wg.Done()
		ch <- 1
		close(ch)
	}

	wg.Add(2)
	go producer()
	go producer()

	for val := range ch {
		slog.Info("successfully received", "value", val)
	}
}

func exitAfter(duration time.Duration) {
	go func() {
		<-time.After(duration)
		slog.Error("timeout exceeded, terminating program.")
		os.Exit(1)
	}()
}
