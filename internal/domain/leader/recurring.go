package leader

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/lloydmeta/tasques/internal/domain/tracing"
)

// InternalRecurringFunction defines a recurring task to run if we have the leader lock
type InternalRecurringFunction struct {
	name string
	// How often to run it
	interval time.Duration
	// What to run
	f func(ctx context.Context, isLeader Checker) error
}

// NewInternalRecurringFunction returns a new recurring task (but doesn't run it)
// Also assumes f is non-nil.
func NewInternalRecurringFunction(name string, interval time.Duration, f func(ctx context.Context, isLeader Checker) error) InternalRecurringFunction {
	return InternalRecurringFunction{name: name, interval: interval, f: f}
}

// InternalRecurringFunctionRunner is a runner of InternalRecurringFunctions.
//
// It is intended to be the centralised place where static, cyclical *internal* background jobs
// are run inside the app.
type InternalRecurringFunctionRunner interface {

	// Start begins the InternalRecurringFunctionRunner loop
	Start()

	// Stop stops the InternalRecurringFunctionRunner loop
	Stop()
}

type impl struct {
	tracer     tracing.Tracer
	functions  []InternalRecurringFunction
	stopped    uint32
	leaderLock Lock
}

// NewInternalRecurringFunctionRunner creates a new InternalRecurringFunctionRunner
func NewInternalRecurringFunctionRunner(tasks []InternalRecurringFunction, tracer tracing.Tracer, leaderLock Lock) InternalRecurringFunctionRunner {
	return &impl{
		functions:  tasks,
		stopped:    1,
		leaderLock: leaderLock,
		tracer:     tracer,
	}
}

// Start begins the InternalRecurringFunctionRunner loop
func (r *impl) Start() {
	atomic.StoreUint32(&r.stopped, 0)
	for _, t := range r.functions {
		go func(task InternalRecurringFunction, shouldRun func() bool, isLeader Checker) {
			for shouldRun() {
				startIterationTime := time.Now().UTC()
				tx := r.tracer.BackgroundTx(task.name)
				ctx := tx.Context()
				err := task.f(ctx, isLeader)
				if err != nil {
					log.Error().Err(err).Msgf("Failed when running task [%s]", task.name)
				}
				tx.End()
				waitTime := task.interval - time.Since(startIterationTime)
				if waitTime > 0 {
					time.Sleep(waitTime)
				}
			}
			log.Info().Msgf("Recurring task ended [%s]", task.name)
		}(t, r.shouldRun, r.leaderLock)
	}
}

// Stop stops the InternalRecurringFunctionRunner loop
func (r *impl) Stop() {
	log.Info().Msg("Stopping recurring functions")
	atomic.StoreUint32(&r.stopped, 1)
}

func (r *impl) shouldRun() bool {
	return atomic.LoadUint32(&r.stopped) == 0
}
