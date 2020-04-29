package gm_plugins

import (
	"sync"

	"gitlab.ziggurat.cn/guomi/gm-plugins/gosm"
	"gitlab.ziggurat.cn/guomi/gm-plugins/primitive"
)

type SmCryptoSuite struct {
	primitive.Sm2Crypto
	primitive.KeysGenerator
}

var (
	smOnce sync.Once
	scs    *SmCryptoSuite
)

func GetSmCryptoSuite() primitive.Context {
	smOnce.Do(func() {
		scs = &SmCryptoSuite{
			Sm2Crypto:     new(gosm.GoSm2),
			KeysGenerator: new(gosm.KeysDerive),
		}
	})

	return scs
}
