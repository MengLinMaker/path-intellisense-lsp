package glsp

import (
	"context"
	"encoding/json"
)

type NotifyFunc func(method string, params any)
type CallFunc func(method string, params any, result any)

type Context struct {
	Method  string
	Params  json.RawMessage
	Notify  NotifyFunc
	Call    CallFunc
	Context context.Context // can be nil
}

type Handler interface {
	Handle(ctx *Context) (result any, validMethod bool, validParams bool, err error)
}
