package closure

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

// FrameSize is the number of frames that FuncStack should trace back from.
const FrameSize = 3

// FuncStack wraps a function value and provides function frames containing the
// caller trace for debugging.
type FuncStack struct {
	Func   reflect.Value
	Frames []uintptr
}

var zeroFuncStack = FuncStack{}

// NewFuncStack creates a new FuncStack. It panics if fn is not a function. The
// given frameSkip is added 2, meaning the first frame from 0 will start from
// the caller of NewFuncStack.
func NewFuncStack(fn interface{}, frameSkip int) FuncStack {
	// Create a reflect.Value from f.  This is called when the returned
	// GClosure runs.
	rf := reflect.ValueOf(fn)

	// Closures can only be created from funcs.
	if rf.Type().Kind() != reflect.Func {
		panic("closure value is not a func")
	}

	frames := make([]uintptr, FrameSize)
	frames = frames[:runtime.Callers(frameSkip+2, frames)]

	return FuncStack{
		Func:   rf,
		Frames: frames,
	}
}

var (
	idleTypeCache    sync.Map
	idleTypeSentinel = struct{}{}
)

// NewIdleFuncStack works akin to NewFuncStack, but it also validates the given
// function type for the correct acceptable signatures for SourceFunc while also
// caching the checks.
func NewIdleFuncStack(fn interface{}, frameSkip int) FuncStack {
	fs := NewFuncStack(fn, frameSkip+1)
	funcType := fs.Func.Type()

	// LoadOrStore will actually ensure that only 1 check is done at a time, but
	// future checks on failed functions may trigger a late panic.
	_, checked := idleTypeCache.LoadOrStore(funcType, idleTypeSentinel)
	if checked {
		return fs
	}

	// Ensure no parameters prematurely.
	if funcType.NumIn() > 0 {
		fs.Panicf("timeout source should have no parameters")
	}

	// Ensure proper return types.
	switch out := funcType.NumOut(); out {
	case 0:
		break
	case 1:
		out0 := funcType.Out(0)
		if out0.Kind() != reflect.Bool {
			fs.Panicf("expected bool in return type, got %v", out0.Kind())
		}
	default:
		fs.Panicf("unexpected return count (expecting 0 or 1): %d", out)
	}

	return fs
}

// IsValid returns true if the given FuncStack is not a zero-value i.e.  valid.
func (fs FuncStack) IsValid() bool {
	return fs.Frames != nil
}

const headerSignature = "closure error: "

// Panicf panics with the given FuncStack printed to standard error.
func (fs FuncStack) Panicf(msgf string, v ...interface{}) {
	msg := strings.Builder{}
	msg.WriteString(headerSignature)
	fmt.Fprintf(&msg, msgf, v...)

	msg.WriteString("\n\nClosure added at:")

	frames := runtime.CallersFrames(fs.Frames)
	for {
		frame, more := frames.Next()
		msg.WriteString("\n\t")
		msg.WriteString(frame.Function)
		msg.WriteString(" at ")
		msg.WriteString(frame.File)
		msg.WriteByte(':')
		msg.WriteString(strconv.Itoa(frame.Line))

		if !more {
			break
		}
	}

	panic(msg.String())
}

// TryRepanic attempts to recover a panic. If successful, it will re-panic with
// the trace, or none if there is already one.
func (fs FuncStack) TryRepanic() {
	panicking := recover()
	if panicking == nil {
		return
	}

	if msg, ok := panicking.(string); ok {
		if strings.HasPrefix(msg, headerSignature) {
			// We can just repanic as-is.
			panic(msg)
		}
	}

	fs.Panicf("unexpected panic caught: %v", panicking)
}
