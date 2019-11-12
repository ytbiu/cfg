package cfg

import (
	"strings"
	"reflect"
)

type Option struct {
	Addr       string
	ConfigPath string
	Result     interface{}

	EncodeType EncType
	hooks      []func()
}

func (o *Option) check() error {
	if strings.TrimSpace(o.Addr) == "" {
		return addrIsEmpty
	}

	if strings.TrimSpace(o.ConfigPath) == ""{
		return configPathIsEmpty
	}

	if o.EncodeType <= ignoreEnc || o.EncodeType >= limitEnc{
		return encodeTypeIsInvalid
	}

	if o.Result == nil || reflect.TypeOf(o.Result).Kind() != reflect.Ptr{
		return resultMustBePtr
	}

	return nil

}

type SetOpt func(opt *Option)

func Addr(addr string) SetOpt {
	return func(opt *Option) {
		opt.Addr = addr
	}
}

func Result(result interface{}) SetOpt {
	return func(opt *Option) {
		opt.Result = result
	}
}

func ConfigPath(path string) SetOpt {
	return func(opt *Option) {
		opt.ConfigPath = path
	}
}

func EncodeType(e EncType) SetOpt {
	return func(opt *Option) {
		opt.EncodeType = e
	}
}

func Hooks(hooks ...func()) SetOpt {
	return func(opt *Option) {
		opt.hooks = hooks
	}
}
