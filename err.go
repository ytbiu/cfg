package cfg

import "errors"

var (
	addrIsEmpty       = errors.New("addr is empty")
	configPathIsEmpty = errors.New("config path is empty")
	resultMustBePtr   = errors.New("result must be pointer")
	encodeTypeIsInvalid = errors.New("encode type is invalid")
)
