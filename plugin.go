package gm_plugins

import (
	"sync"

	"github.com/zhigui-projects/gm-plugins/gosm"
	"github.com/zhigui-projects/gm-plugins/primitive"
)

type SmCryptoSuite struct {
	primitive.Sm2Crypto
	primitive.Sm3Crypro
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
			Sm3Crypro:     new(gosm.GoSm3),
			KeysGenerator: new(gosm.KeysDerive),
		}
	})

	return scs
}
