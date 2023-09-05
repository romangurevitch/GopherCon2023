package simulator

import (
	"context"
	"log/slog"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// Event encapsulates an I/O event with its creation timestamp and an interval indicator.
type Event struct {
	CreatedAt time.Time
	Interval  int
}

// EventResult combines an Event with a HandledAt timestamp to indicate when the event was processed.
type EventResult struct {
	Event
	HandledAt time.Time
}

// Config defines the parameters to configure the Simulator's behavior.
type Config struct {
	NumberOfChannels int           // Number of channels to simulate.
	MinInterval      time.Duration // Minimum time between emitted events per channel.
	MaxInterval      time.Duration // Maximum time between emitted events per channel.
	UpdateRate       time.Duration // Rate at which the interval time is updated.
	IntervalStep     time.Duration // Amount by which the interval time is decreased at each update.
	MaxJitter        time.Duration // Maximum random jitter added to emission intervals.
}

// Simulator manages a simulation environment for emitting events across multiple channels.
type Simulator struct {
	channels     []chan Event       // Channels used to emit events.
	config       Config             // Configuration parameters for the simulator.
	stats        stats              // Internal statistics and estimates based on the configuration.
	ctx          context.Context    // Context to control the lifecycle of the simulation.
	cancel       context.CancelFunc // Cancel function to stop the simulation.
	waitGroup    *sync.WaitGroup    // WaitGroup to synchronize the completion of goroutines.
	eventCounter atomic.Int64       // Counter for the total number of emitted events.
}

// NewSimulator initializes a new Simulator instance with the given configuration and context.
func NewSimulator(ctx context.Context, config Config) *Simulator {
	ctx, cancel := context.WithCancel(ctx)
	s := &Simulator{
		channels:  make([]chan Event, config.NumberOfChannels),
		config:    config,
		stats:     calculateStats(config),
		ctx:       ctx,
		cancel:    cancel,
		waitGroup: &sync.WaitGroup{},
	}

	// Prepare buffered channels based on the expected number of events per channel.
	for i := range s.channels {
		s.channels[i] = make(chan Event, s.stats.sumEventsPerChannel)
	}

	return s
}

// simulateChannel runs a goroutine that emits events on the given channel at dynamically changing intervals.
func (s *Simulator) simulateChannel(ch chan<- Event) {
	defer s.waitGroup.Done()
	defer close(ch)

	currentInterval := s.config.MaxInterval
	ticker := time.NewTicker(currentInterval)
	defer ticker.Stop()

	incrementTicker := time.NewTicker(s.config.UpdateRate)
	defer incrementTicker.Stop()

	intervalCount := 0

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-incrementTicker.C:
			currentInterval -= s.config.IntervalStep
			intervalCount++
			if currentInterval < s.config.MinInterval {
				return
			}

			ticker.Reset(currentInterval)
		case <-ticker.C:
			s.emitEvent(ch, intervalCount)

			jitter := time.Duration(rand.Int63n(int64(s.config.MaxJitter)))
			ticker.Reset(currentInterval + jitter)
		}
	}
}

// emitEvent sends a new Event to the specified channel and increments the event counter.
func (s *Simulator) emitEvent(ch chan<- Event, intervalCount int) {
	event := Event{CreatedAt: time.Now(), Interval: intervalCount}
	select {
	case <-s.ctx.Done():
		return
	case ch <- event:
		s.eventCounter.Add(1)
	}
}

// Start launches the simulation by starting goroutines for each channel to emit events.
func (s *Simulator) Start() {
	for _, ch := range s.channels {
		s.waitGroup.Add(1)
		go s.simulateChannel(ch)
	}
}

// Wait blocks until all event-emitting goroutines have finished.
func (s *Simulator) Wait() {
	s.waitGroup.Wait()
}

// EventCount returns the total number of events emitted during the simulation.
func (s *Simulator) EventCount() int64 {
	return s.eventCounter.Load()
}

// GetChannels provides access to the channels used in the simulation.
func (s *Simulator) GetChannels() []chan Event {
	return s.channels
}

// PrintConfig outputs the configuration and estimated statistics of the simulation to the log.
func (s *Simulator) PrintConfig() {
	slog.Info("Simulator", "number of channels", s.config.NumberOfChannels)
	slog.Info("Simulator", "total intervals", s.stats.totalIntervals)
	slog.Info("Simulator", "min samples per interval", s.stats.minSamplesPerInterval, "starting at", s.config.MaxInterval)
	slog.Info("Simulator", "max samples per interval", s.stats.maxSamplesPerInterval, "ending at", s.config.MinInterval)
	slog.Info("Simulator", "estimated sum of all events per channel", s.stats.sumEventsPerChannel)
	slog.Info("Simulator", "estimated total events", s.stats.sumAllEvents)
	slog.Info("Simulator", "estimated runtime", s.stats.estimateRuntime)
}

// stats holds calculated statistics based on the simulator configuration.
type stats struct {
	totalIntervals        int           // Total number of interval updates.
	minSamplesPerInterval int           // Minimum number of samples per update interval.
	maxSamplesPerInterval int           // Maximum number of samples per update interval.
	sumEventsPerChannel   int           // Total number of events expected per channel.
	sumAllEvents          int           // Total number of events expected across all channels.
	estimateRuntime       time.Duration // Estimated duration of the simulation.
}

// calculateStats computes the statistics based on the simulator configuration.
func calculateStats(config Config) stats {
	intervalCount := int((config.MaxInterval - config.MinInterval) / config.IntervalStep)
	minSamples := int(config.UpdateRate / config.MaxInterval)
	maxSamples := int(config.UpdateRate / config.MinInterval)

	return stats{
		totalIntervals:        intervalCount,
		minSamplesPerInterval: minSamples,
		maxSamplesPerInterval: maxSamples,
		sumEventsPerChannel:   intervalCount * (minSamples + maxSamples) / 2,
		sumAllEvents:          intervalCount * (minSamples + maxSamples) / 2 * config.NumberOfChannels,
		estimateRuntime:       time.Duration(intervalCount) * config.UpdateRate,
	}
}
