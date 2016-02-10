package iface


type Surface interface {
    CopyPage()
    CreateForRectangle(float64, float64, float64, float64) Surface
    CreateSimilar(Content, int, int) Surface
    Flush()
    GetDeviceOffset() (float64, float64)
    GetFallbackResolution() (float64, float64)
    GetMimeData(MimeType) []byte
    GetType() SurfaceType
    HasShowTextGlyphs() bool
    MarkDirty()
    MarkDirtyRectangle(int, int, int, int)
    SetDeviceOffset(float64, float64)
    SetFallbackResolution(float64, float64)
    ShowPage()
    Status() Status
} // end of Surface

func AssertSurface(_ Surface) {}
