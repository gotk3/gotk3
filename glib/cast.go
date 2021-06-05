package glib

type WrapFn interface{}

var WrapMap = map[string]WrapFn{
	"GMenu": wrapMenuModel,
}