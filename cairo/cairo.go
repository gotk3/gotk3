package cairo

type Cairo interface {
	Create(Surface) Context
	NewSurface(uintptr, bool) Surface
	NewSurfaceFromPNG(string) (Surface, error)
	StatusToString(Status) string
} // end of Cairo

func AssertCairo(_ Cairo) {}
