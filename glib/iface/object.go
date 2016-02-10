package iface


type Object interface {
    Connect(string, interface{}, ...interface{}) (SignalHandle, error)
    ConnectAfter(string, interface{}, ...interface{}) (SignalHandle, error)
    Emit(string, ...interface{}) (interface{}, error)
    ForceFloating()
    GetProperty(string) (interface{}, error)
    GetPropertyType(string) (Type, error)
    HandlerBlock(SignalHandle)
    HandlerDisconnect(SignalHandle)
    HandlerUnblock(SignalHandle)
    IsA(Type) bool
    IsFloating() bool
    Ref()
    RefSink()
    Set(string, interface{}) error
    SetProperty(string, interface{}) error
    StopEmission(string)
    TypeFromInstance() Type
    Unref()
} // end of Object

func AssertObject(_ Object) {}
