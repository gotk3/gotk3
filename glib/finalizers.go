package glib

// Finalizer is a function that when called will finalize an object
type Finalizer func()

// FinalizerStrategy will be called by every runtime finalizer in gotk3
// The simple version will just call the finalizer given as an argument
// but in larger programs this might cause problems with the UI thread.
// The FinalizerStrategy function will always be called in the goroutine that
// `runtime.SetFinalizer` uses. It is a `var` to explicitly allow clients to
// change the strategy to something more advanced.
var FinalizerStrategy = func(f Finalizer) {
	f()
}
