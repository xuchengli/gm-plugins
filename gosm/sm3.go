package gosm

import (
	"hash"

	"gitlab.ziggurat.cn/guomi/gm-go/sm3"
)

type GoSm3 struct{}

func (gs *GoSm3) NewSm3() hash.Hash {
	return sm3.New()
}
