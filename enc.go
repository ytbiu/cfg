package cfg

import (
	"github.com/micro/go-micro/config/encoder"
	"github.com/micro/go-micro/config/encoder/json"
	"github.com/micro/go-micro/config/encoder/yaml"
)

type EncType uint8

const (
	ignoreEnc = iota

	YamlEnc
	JsonEnc

	limitEnc
)

func (e EncType) NewEncoder() encoder.Encoder {
	switch e {
	case YamlEnc:
		return yaml.NewEncoder()
	case JsonEnc:
		return json.NewEncoder()
	}
	return nil
}
