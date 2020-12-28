// +build go1.15

package closure

import "unsafe"

func getAndDeleteClosure(closure unsafe.Pointer) FuncStack {
	v, ok := closures.LoadAndDelete(closure)
	if ok {
		return v.(FuncStack)
	}
	return zeroFuncStack
}
