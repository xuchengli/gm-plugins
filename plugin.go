package gm_plugins

import (
	"sync"

	"gitlab.ziggurat.cn/guomi/gm-common/sm"
	"gitlab.ziggurat.cn/guomi/gm-plugins/gosm"
)

type SmCryptoSuite struct {
	sm.Sm2Crypto
	sm.KeysGenerator
}

var (
	smOnce sync.Once
	scs    *SmCryptoSuite
)

func GetSmCryptoSuite() sm.Context {
	smOnce.Do(func() {
		scs = &SmCryptoSuite{
			Sm2Crypto:     new(gosm.GoSm2),
			KeysGenerator: new(gosm.KeysDerive),
		}
	})

	return scs
}
