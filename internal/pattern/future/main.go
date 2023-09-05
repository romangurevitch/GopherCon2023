package main

import (
	"context"
	"log/slog"
	"time"

	"github.com/mtslzr/pokeapi-go"
	"github.com/mtslzr/pokeapi-go/structs"
)

// Result type represents a computation result.
type Result[T any] struct {
	Value T
	Err   error
}

// Future type represents a future value.
type Future[T any] struct {
	result chan Result[T] // result is a channel that will contain the result.
}

// ProcessFunc defines a function type for processing a value of type T to produce a value of type U, in a context-aware manner.
type ProcessFunc[T any] func(context.Context) (T, error)

// NewFuture creates a new Future.
func NewFuture[T any](ctx context.Context, processFunc ProcessFunc[T]) *Future[T] {
	f := &Future[T]{result: make(chan Result[T], 1)} // Buffered channel to prevent blocking.
	go func() {
		defer close(f.result)
		select {
		case <-ctx.Done():
			f.result <- Result[T]{Err: ctx.Err()} // Send context error if it was canceled.
		default:
			value, err := processFunc(ctx)
			f.result <- Result[T]{Value: value, Err: err} // Send processFunc result.
		}
	}()
	return f
}

// Result retrieves the result of the computation.
func (f *Future[T]) Result() Result[T] {
	return <-f.result // This will block until the result is ready.
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // Ensure all resources are cleaned up

	future := NewFuture(ctx, func(ctx context.Context) (structs.Pokemon, error) {
		return pokeapi.Pokemon("pikachu")
	})

	// Optionally, do some other work here while waiting for the future result...

	// Now wait for the result:
	result := future.Result()
	if result.Err != nil {
		slog.Error("Error fetching Pokémon details", "error", result.Err)
		return
	}
	slog.Info("Fetched Pokémon details", "pokemonName", result.Value.Name)
}
