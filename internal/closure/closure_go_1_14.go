// +build !go1.15

package closure

import "unsafe"

func getAndDeleteClosure(closure unsafe.Pointer) FuncStack {
	v, ok := closures.Load(closure)
	if ok {
		closures.Delete(closure)
		return v.(FuncStack)
	}
	return zeroFuncStack
}
