// +build !gtk_3_6
// Supports building with gtk 3.8+

/*
 * Copyright (c) 2013-2014 Conformal Systems <info@conformal.com>
 *
 * This file originated from: http://opensource.conformal.com/
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package gdk

// #include <gdk/gdk.h>
// #include "gdk_since_3_8.go.h"
import "C"
import (
	"unsafe"

	"github.com/gotk3/gotk3/glib"
)

func init() {

	tm := []glib.TypeMarshaler{
		// Enums
		{glib.Type(C.gdk_frame_clock_phase_get_type()), marshalClockPhase},

		// Objects/Interfaces
		{glib.Type(C.gdk_frame_clock_get_type()), marshalFrameClock},
		{glib.Type(C.gdk_frame_timings_get_type()), marshalFrameTimings},
	}

	glib.RegisterGValueMarshalers(tm)
}

// ClockPhase is a representation of GDK's GdkFrameClockPhase.
type ClockPhase int

const (
	PHASE_NONE          ClockPhase = C.GDK_FRAME_CLOCK_PHASE_NONE
	PHASE_FLUSH_EVENTS  ClockPhase = C.GDK_FRAME_CLOCK_PHASE_FLUSH_EVENTS
	PHASE_BEFORE_PAINT  ClockPhase = C.GDK_FRAME_CLOCK_PHASE_BEFORE_PAINT
	PHASE_UPDATE        ClockPhase = C.GDK_FRAME_CLOCK_PHASE_UPDATE
	PHASE_LAYOUT        ClockPhase = C.GDK_FRAME_CLOCK_PHASE_LAYOUT
	PHASE_PAINT         ClockPhase = C.GDK_FRAME_CLOCK_PHASE_PAINT
	PHASE_RESUME_EVENTS ClockPhase = C.GDK_FRAME_CLOCK_PHASE_RESUME_EVENTS
	PHASE_AFTER_PAINT   ClockPhase = C.GDK_FRAME_CLOCK_PHASE_AFTER_PAINT
)

func marshalClockPhase(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return ClockPhase(c), nil
}

/*
 * GdkFrameClock
 */

// FrameClock is a representation of GDK's GdkFrameClock.
type FrameClock struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkFrameClock.
func (v *FrameClock) native() *C.GdkFrameClock {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkFrameClock(p)
}

// Native returns a pointer to the underlying GdkFrameClock.
func (v *FrameClock) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalFrameClock(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return WrapFrameClock(unsafe.Pointer(c)), nil
}

func WrapFrameClock(ptr unsafe.Pointer) *FrameClock {
	obj := &glib.Object{glib.ToGObject(ptr)}
	return &FrameClock{obj}
}

// BeginUpdating is a wrapper around gdk_frame_clock_begin_updating().
func (v *FrameClock) BeginUpdating() {
	C.gdk_frame_clock_begin_updating(v.native())
}

// EndUpdating is a wrapper around gdk_frame_clock_end_updating().
func (v *FrameClock) EndUpdating() {
	C.gdk_frame_clock_end_updating(v.native())
}

// GetFrameTime is a wrapper around gdk_frame_clock_get_frame_time().
func (v *FrameClock) GetFrameTime() int64 {
	return int64(C.gdk_frame_clock_get_frame_time(v.native()))
}

// GetFrameCounter is a wrapper around gdk_frame_clock_get_frame_counter().
func (v *FrameClock) GetFrameCounter() int64 {
	return int64(C.gdk_frame_clock_get_frame_counter(v.native()))
}

// GetHistoryStart is a wrapper around gdk_frame_clock_get_history_start().
func (v *FrameClock) GetHistoryStart() int64 {
	return int64(C.gdk_frame_clock_get_history_start(v.native()))
}

// GetTimings is a wrapper around gdk_frame_clock_get_timings().
func (v *FrameClock) GetTimings(frameCounter int64) (*FrameTimings, error) {
	c := C.gdk_frame_clock_get_timings(v.native(), C.gint64(frameCounter))
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapFrameTimings(unsafe.Pointer(c)), nil
}

// GetCurrentTimings is a wrapper around dk_frame_clock_get_current_timings().
func (v *FrameClock) GetCurrentTimings() (*FrameTimings, error) {
	c := C.gdk_frame_clock_get_current_timings(v.native())
	if c == nil {
		return nil, nilPtrErr
	}
	return wrapFrameTimings(unsafe.Pointer(c)), nil
}

// GetRefreshInfo is a wrapper around gdk_frame_clock_get_refresh_info().
func (v *FrameClock) GetRefreshInfo(baseTime int64) (int64, int64) {
	var cr, cp (*C.gint64)
	defer C.free(unsafe.Pointer(cr))
	defer C.free(unsafe.Pointer(cp))
	b := C.gint64(baseTime)

	C.gdk_frame_clock_get_refresh_info(v.native(), b, cr, cp)
	r, p := int64(*cr), int64(*cp)
	return r, p
}

// RequestPhase is a wrapper around gdk_frame_clock_request_phase().
func (v *FrameClock) RequestPhase(phase ClockPhase) {
	C.gdk_frame_clock_request_phase(v.native(), C.GdkFrameClockPhase(phase))
}

/*
 * GdkFrameTimings
 */

// FrameTimings is a representation of GDK's GdkFrameTimings.
type FrameTimings struct {
	*glib.Object
}

// native returns a pointer to the underlying GdkFrameTimings.
func (v *FrameTimings) native() *C.GdkFrameTimings {
	if v == nil || v.GObject == nil {
		return nil
	}
	p := unsafe.Pointer(v.GObject)
	return C.toGdkFrameTimings(p)
}

// Native returns a pointer to the underlying GdkFrameTimings.
func (v *FrameTimings) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func wrapFrameTimings(ptr unsafe.Pointer) *FrameTimings {
	obj := &glib.Object{glib.ToGObject(ptr)}
	return &FrameTimings{obj}
}

func marshalFrameTimings(p uintptr) (interface{}, error) {
	c := C.g_value_get_object((*C.GValue)(unsafe.Pointer(p)))
	return wrapFrameTimings(unsafe.Pointer(c)), nil
}

// Ref is a wrapper around gdk_frame_timings_ref().
func (v *FrameTimings) Ref() {
	c := C.gdk_frame_timings_ref(v.native())
	v = wrapFrameTimings(unsafe.Pointer(c))
}

// Unref is a wrapper around gdk_frame_timings_unref().
func (v *FrameTimings) Unref() {
	C.gdk_frame_timings_unref(v.native())
}

// GetFrameCounter is a wrapper around gdk_frame_timings_get_frame_counter().
func (v *FrameTimings) GetFrameCounter() int64 {
	return int64(C.gdk_frame_timings_get_frame_counter(v.native()))
}

// GetComplete is a wrapper around gdk_frame_timings_get_complete().
func (v *FrameTimings) GetComplete() bool {
	return gobool(C.gdk_frame_timings_get_complete(v.native()))
}

// GetFrameTime is a wrapper around gdk_frame_timings_get_frame_time().
func (v *FrameTimings) GetFrameTime() int64 {
	return int64(C.gdk_frame_timings_get_frame_time(v.native()))
}

// GetPresentationTime is a wrapper around gdk_frame_timings_get_presentation_time().
func (v *FrameTimings) GetPresentationTime() int64 {
	return int64(C.gdk_frame_timings_get_presentation_time(v.native()))
}

// GetRefreshInterval is a wrapper around gdk_frame_timings_get_refresh_interval().
func (v *FrameTimings) GetRefreshInterval() int64 {
	return int64(C.gdk_frame_timings_get_refresh_interval(v.native()))
}

// GetPredictedPresentationTime is a wrapper around gdk_frame_timings_get_predicted_presentation_time().
func (v *FrameTimings) GetPredictedPresentationTime() int64 {
	return int64(C.gdk_frame_timings_get_predicted_presentation_time(v.native()))
}
