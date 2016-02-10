package iface


type Context interface {
    Arc(float64, float64, float64, float64, float64)
    ArcNegative(float64, float64, float64, float64, float64)
    Clip()
    ClipExtents() (float64, float64, float64, float64)
    ClipPreserve()
    ClosePath()
    CopyPage()
    CurveTo(float64, float64, float64, float64, float64, float64)
    Fill()
    FillExtents() (float64, float64, float64, float64)
    FillPreserve()
    GetAntialias() Antialias
    GetCurrentPoint() (float64, float64)
    GetDash() ([]float64, float64)
    GetDashCount() int
    GetFillRule() FillRule
    GetGroupTarget() Surface
    GetLineCap() LineCap
    GetLineJoin() LineJoin
    GetLineWidth() float64
    GetMiterLimit() float64
    GetOperator() Operator
    GetTarget() Surface
    GetTolerance() float64
    InClip(float64, float64) bool
    InFill(float64, float64) bool
    InStroke(float64, float64) bool
    LineTo(float64, float64)
    MaskSurface(Surface, float64, float64)
    MoveTo(float64, float64)
    NewPath()
    Paint()
    PaintWithAlpha(float64)
    PopGroupToSource()
    PushGroup()
    PushGroupWithContent(Content)
    Rectangle(float64, float64, float64, float64)
    ResetClip()
    Restore()
    Save()
    SetAntialias(Antialias)
    SetDash([]float64, float64)
    SetFillRule(FillRule)
    SetLineCap(LineCap)
    SetLineJoin(LineJoin)
    SetLineWidth(float64)
    SetMiterLimit(float64)
    SetOperator(Operator)
    SetSourceRGB(float64, float64, float64)
    SetSourceRGBA(float64, float64, float64, float64)
    SetSourceSurface(Surface, float64, float64)
    SetTolerance(float64)
    ShowPage()
    Status() Status
    Stroke()
    StrokeExtents() (float64, float64, float64, float64)
    StrokePreserve()
} // end of Context

func AssertContext(_ Context) {}
