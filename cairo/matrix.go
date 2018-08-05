package cairo

// #cgo pkg-config: cairo cairo-gobject
// #include <stdlib.h>
// #include <cairo.h>
// #include <cairo-gobject.h>
import "C"

import (
	"unsafe"
)

// Matrix struct
type Matrix struct {
	Xx, Yx float64
	Xy, Yy float64
	X0, Y0 float64
}

// NewMatrix creates a new identiy matrix
func NewMatrix(xx, yx, xy, yy, x0, y0 float64) *Matrix {
	return &Matrix{
		Xx: xx,
		Yx: yx,
		Xy: xy,
		Yy: yy,
		X0: x0,
		Y0: y0,
	}
}

// Native returns native c pointer to a matrix
func (m *Matrix) Native() *C.cairo_matrix_t {
	return (*C.cairo_matrix_t)(unsafe.Pointer(m))
}

// InitIdentity initializes this matrix to identity matrix
func (m *Matrix) InitIdentity() {
	C.cairo_matrix_init_identity(m.Native())
}

// InitTranslate initializes a matrix with the given translation
func (m *Matrix) InitTranslate(tx, ty float64) {
	C.cairo_matrix_init_translate(m.Native(), C.double(tx), C.double(ty))
}

// InitScale initializes a matrix with the give scale
func (m *Matrix) InitScale(sx, sy float64) {
	C.cairo_matrix_init_scale(m.Native(), C.double(sx), C.double(sy))
}

// InitRotate initializes a matrix with the given rotation
func (m *Matrix) InitRotate(radians float64) {
	C.cairo_matrix_init_rotate(m.Native(), C.double(radians))
}

// Translate translates a matrix by the given amount
func (m *Matrix) Translate(tx, ty float64) {
	C.cairo_matrix_translate(m.Native(), C.double(tx), C.double(ty))
}

// Scale scales the matrix by the given amounts
func (m *Matrix) Scale(sx, sy float64) {
	C.cairo_matrix_scale(m.Native(), C.double(sx), C.double(sy))
}

// Rotate rotates the matrix by the given amount
func (m *Matrix) Rotate(radians float64) {
	C.cairo_matrix_rotate(m.Native(), C.double(radians))
}

// Invert inverts the matrix
func (m *Matrix) Invert() {
	C.cairo_matrix_invert(m.Native())
}

// Multiply multiplies the matrix by another matrix
func (m *Matrix) Multiply(a, b Matrix) {
	C.cairo_matrix_multiply(m.Native(), a.Native(), b.Native())
}

// TransformDistance ...
func (m *Matrix) TransformDistance(dx, dy float64) (float64, float64) {
	C.cairo_matrix_transform_distance(m.Native(),
		(*C.double)(unsafe.Pointer(&dx)), (*C.double)(unsafe.Pointer(&dy)))
	return dx, dy
}

// TransformPoint ...
func (m *Matrix) TransformPoint(x, y float64) (float64, float64) {
	C.cairo_matrix_transform_point(m.Native(),
		(*C.double)(unsafe.Pointer(&x)), (*C.double)(unsafe.Pointer(&y)))
	return x, y
}
