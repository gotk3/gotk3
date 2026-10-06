package glib_test

import (
	"runtime"
	"sync"
	"testing"
	"time"

	"github.com/gotk3/gotk3/glib"
)

// growStack recurses depth times, so that the goroutine's stack grows and is
// then mostly unused, which makes the garbage collector shrink it.
//
//go:noinline
func growStack(depth int) {
	var pad [64]byte
	if depth > 0 {
		growStack(depth - 1)
	}
	_ = pad
}

// TestSourceAddStackShrink adds and removes idle and timeout sources while the
// garbage collector runs continuously and shrinks the adding goroutines'
// stacks. Shrinking a stack copies it, and the copy aborts the program with
// "invalid pointer found on stack" if a callback ID, a small integer, sits in
// one of its pointer slots.
func TestSourceAddStackShrink(t *testing.T) {
	duration := 5 * time.Second
	if testing.Short() {
		duration = time.Second
	}
	deadline := time.Now().Add(duration)

	var wg sync.WaitGroup
	for i := 0; i < runtime.GOMAXPROCS(0); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for time.Now().Before(deadline) {
				growStack(1024)
				for j := 0; j < 100; j++ {
					glib.SourceRemove(glib.IdleAdd(func() {}))
					glib.SourceRemove(glib.TimeoutAdd(1000, func() {}))
				}
			}
		}()
	}

	stop := make(chan struct{})
	gcDone := make(chan struct{})
	go func() {
		defer close(gcDone)
		for {
			select {
			case <-stop:
				return
			default:
				runtime.GC()
			}
		}
	}()

	wg.Wait()
	close(stop)
	<-gcDone
}
